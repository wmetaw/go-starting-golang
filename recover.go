package main

import "fmt"

func testRecover(src interface{}) {
	defer func() {

		// recoverはinterface{}の値を戻し、その値がnilでなければpanicが実行されたと判断できる
		// 変数xはpanicに渡されたinterface{}
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Println("panic: unknown")
			}
			fmt.Println(x)
		}
	}()
	panic(src)
	return
}

func main() {
	testRecover(128)
	testRecover("hoge")
	testRecover([...]int{1, 2, 3})
}
