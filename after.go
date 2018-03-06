package main

import (
	"fmt"
	"time"
)

func main() {
	// time.After 開始から1秒後に時刻を表示
	ch := time.After(1 * time.Second)
	fmt.Println(<-ch)

	// goroutineで何らかの処理
	c := make(chan string)
	go func(chan string) {
		time.Sleep(3 * time.Second)
		c <- "goroutine!"
	}(c)

	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout...")
	}
}
