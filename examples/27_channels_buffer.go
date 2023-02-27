package examples

import "fmt"

// BufferedChannel allows data to be sent into channel (chan -> ) wihout a corresponding receive (<- chan)

func GoChannelsBuffer() {
	message := make(chan string, 2)

	message <- "Hello"
	message <- "World"
	// message <- "There"
	fmt.Println(<-message)
	fmt.Println(<-message)
	// fmt.Println(<-message)
}
