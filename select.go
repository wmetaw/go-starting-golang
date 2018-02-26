package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// 整数を増分させつつch1へ送信
	go func() {
		for {
			i := <-ch1
			ch2 <- i * 2
		}
	}()

	// ch3から受信した整数を出力
	go func() {
		for {
			i := <-ch2
			ch3 <- i - 1
		}
	}()

	n := 1
LOOP:
	for {
		select {
		case ch1 <- n:
			n++
		case i := <-ch3:
			fmt.Println("received", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}
