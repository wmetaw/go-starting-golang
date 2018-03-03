package main

import "fmt"

type T struct {
	Id   int
	Name string
}

func (t *T) String() string {
	return fmt.Sprintf("<<%d,%s>>", t.Id, t.Name)
}

func main() {
	t := &T{Id: 10, Name: "Taro"}

	// default => &{10 Taro}
	// Custom  => <<10,Taro>>
	fmt.Println(t)
}
