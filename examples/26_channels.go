package examples

import (
	"fmt"
	"strconv"
	"time"
)

type Msg struct {
	listenerId string
	message    string
}

var channel1 = make(chan string)

func sendMessageToChannel1(msg string) {
	channel1 <- msg
}

var logicChannelRequest = make(chan Msg)
var logicChannelResponse = make(chan Msg)

func Listener(id string, delayed int, p chan string) {
	for i := 0; i < delayed; i++ {
		time.Sleep(time.Second)
		comingMsg := <-p
		fmt.Println("got " + comingMsg + "; now pass to Logic")
		logicChannelRequest <- Msg{
			listenerId: id,
			message:    comingMsg,
		}

		resp := <-logicChannelResponse
		if resp.listenerId == id {
			fmt.Printf("++ %s response: %s\n", id, resp.message)
		}
	}
	close(p)
}

func Logic(delayed int) {
	hstr := make(map[string]int)
	for i := 0; i < delayed; i++ {
		time.Sleep(time.Second)
		comingMsg := <-logicChannelRequest
		listenerId := comingMsg.listenerId
		hstr[listenerId] += 1
		fmt.Printf("-- %d(th) message from %s; with content: %s\n", hstr[listenerId], listenerId, comingMsg.message)
		logicChannelResponse <- Msg{
			listenerId: listenerId,
			message:    "total message: " + strconv.Itoa(hstr[listenerId]),
		}
	}
}

func GoChannels() {
	go sendMessageToChannel1("hello world")
	// local variable in main goroutine
	msg := <-channel1
	fmt.Println("msg in main goroutine: ", msg)

	// create 5 listener with 5 channels
	c := make(map[int]chan string)
	for i := 0; i < 5; i++ {
		c[i] = make(chan string)
		go Listener("l"+strconv.Itoa(i), 5, c[i])
	}
	go Logic(20)
	// go Logic(20)
	// let's shout multiple thing into listeners
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		msg := "hey_" + strconv.Itoa(i)
		fmt.Println("let's shout: " + msg)
		c[i%5] <- msg
	}
}
