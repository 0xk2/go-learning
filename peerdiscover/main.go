package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

var peer = make(chan string)

func serve(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	log.Println("connect ", c.LocalAddr().String())
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		peer <- string(message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func startServer(p string) {
	log.Println("Start server at " + p)
	http.HandleFunc("/ws", serve)
	log.Fatal(http.ListenAndServe(":"+p, nil))
}

func startClient(remoteServer string, selfServer string) {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: ":" + remoteServer, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	c.WriteJSON(selfServer)
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		// case t := <-ticker.C:
		// 	err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		// 	if err != nil {
		// 		log.Println("write:", err)
		// 		return
		// 	}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func clientMaster(selfServer string) {
	log.Println("clientMaster ", selfServer)
	for {
		remotePort := <-peer
		go startClient(remotePort, selfServer)
	}
}

func main() {
	serverPort := os.Args[1]
	defaultPeer := ""
	if len(os.Args) > 2 {
		defaultPeer = os.Args[2]
	}
	log.Println(defaultPeer)
	// connect to peer from configuration (port)
	if defaultPeer != "" {
		// connect
		go startClient(defaultPeer, serverPort)
		go clientMaster(serverPort)
	}
	// open server at a port
	flag.Parse()
	log.SetFlags(0)
	go startServer(serverPort)
	for {
	}
}
