package main

import "fmt"

func later() func(string) string {

	// 一つ前の文字列を保存するための変数
	var store string

	// 引数に文字列を取り、文字列を返す関数を返す
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {

	f := later()

	fmt.Println(f("Golang"))
	fmt.Println(f("is"))
	fmt.Println(f("awesome"))
}
