package basic_tutorial

import (
	"fmt"
	"time"
)

func workersWithChannels(id int, ch chan string) {
	time.Sleep(time.Second)
	fmt.Println("acasaaa")
	ch <- fmt.Sprintf("Worker %d done", id)
}

func CallChannels() {
	newCh := make(chan string)

	for i := 1; i <= 3; i++ {
		go workersWithChannels(i, newCh)
	}

	for i := 1; i <= 3; i++ {
		msg := <-newCh
		fmt.Println(msg)
	}
}
