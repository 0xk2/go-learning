package examples

import (
	"fmt"
	"strconv"
	"time"
)

func worker(_done chan bool, str string) {
	fmt.Printf("making %s ...", str)
	time.Sleep(2 * time.Second)
	fmt.Printf("done\n")
	_done <- true
}

func GoChannelsSync() {
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(done, strconv.Itoa(i+1))
		<-done
	}
}
