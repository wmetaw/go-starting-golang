package main

import "fmt"

/* チャネルの型

var (
 ch0 chan int    送受信
 ch1 <-chan int  受信専用
 ch2 chan<- int  送信専用
)

ch1 = ch0  // OK
ch2 = ch0  // OK
ch0 = ch1  // NG
ch1 = ch2  // NG
*/

// 受信専用チャネルを引数に受け取る
func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go receiver(ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
}

/* 以下の場合はデッドロック
mainスレッドのgoroutineは他goroutineから受信を待ち続け、ランタイムパニックとなる

func main() {
	ch := make(chan int)
	fmt.Println(<-ch)
}
*/
