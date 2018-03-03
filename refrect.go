package main

import (
	"fmt"
	"reflect"
)

type T struct {
	N int `json:"n"`
	s string
}

func main() {

	// オブジェクトの生成
	typ := reflect.TypeOf(100)
	obj := reflect.New(typ)
	fmt.Println(obj)

	// 変数に値を設定（ポインタ必須
	n := 100
	r := reflect.ValueOf(&n)
	r.Elem().SetInt(200)
	fmt.Println(n)

	// 構造体の値を取得
	t := T{N: 100, s: "hoge"}
	tp := reflect.ValueOf(&t)
	tp.Elem().FieldByName("N").SetInt(300)
	fmt.Println(t)
}
