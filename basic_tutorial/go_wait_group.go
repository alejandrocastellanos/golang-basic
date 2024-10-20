package basic_tutorial

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	fmt.Printf("worker %d starging\n", id)
	time.Sleep(time.Second)
	time := time.Since(start)
	fmt.Printf("%v, worker %d DONE\n", time, id)
}

func CallWorker() {
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(1, &wg)

	wg.Add(1)
	go worker(2, &wg)

	wg.Wait()
	fmt.Println("All workers done")
}
