package main

import (
	"fmt"
	"time"
)

func main() {

	ch := time.Tick(1 * time.Second) // <-chan Time
	for {
		t := <-ch // tはtime.Time型
		fmt.Println(t)
	}
}
