package basic_tutorial

import (
	"fmt"
	"time"
)

func SayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func SayWorld() {
	for i := 0; i < 5; i++ {
		fmt.Println("World")
		time.Sleep(100 * time.Millisecond)
	}
}

func GorutinesBasic() {
	go SayHello()
	go SayWorld()
	time.Sleep(1 * time.Second)
}
