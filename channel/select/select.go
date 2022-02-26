package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	var i int
	go func() {
		for true {
			out <- i
			i++
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	c1, c2 := generator(), generator()
	w := createWorker(0)
	var tempArr []int
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for true {
		var activeWorker chan<- int
		var activeVal int
		if len(tempArr) > 0 {
			activeWorker = w
			activeVal = tempArr[0]
		}
		select {
		case n := <-c1:
			tempArr = append(tempArr, n)
		case n := <-c2:
			tempArr = append(tempArr, n)
		case activeWorker <- activeVal:
			tempArr = tempArr[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out")
		case <-tick:
			fmt.Printf("tempArr's length is %d,%v\n", len(tempArr), tempArr)
		case <-after:
			fmt.Printf("Bye")
			return
		}
	}
}
