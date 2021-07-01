package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)

	b := make(chan int)

	for i := 0; i < 50; i++ {
		go func(i int) {
			a <- i
			fmt.Println("insert into a", i)
			//defer wg.Done()
		}(i)
	}

	go func() {
		for {
			b <- <-a
			fmt.Println("read from a")
			time.Sleep(time.Second)
		}
	}()

	//time.Sleep(time.Second * 5)
	time.Sleep(time.Second * 600)
}
