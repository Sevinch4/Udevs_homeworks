package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go task1(ch)
	fmt.Println(<-ch)

	ch2 := make(chan int)
	task2(ch2)
	go func() {
		for {
			fmt.Println(<-ch2)
		}
	}()
	time.Sleep(time.Second)
}

func task1(ch chan<- int) {
	n := rand.Intn(100) + 1
	ch <- n
}
func task2(ch chan<- int) {
	n := rand.Intn(10) + 1
	fmt.Println("rand num:", n)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()

}
