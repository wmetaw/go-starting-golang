package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 一般的な現在時刻をseed値に入れる方法
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())

	// 0 ~ 99までの乱数
	n := rand.Intn(100)
	fmt.Println(n)

	// 暗号化的擬似乱数
	// "crypto/rand"から binary.Readでint64の値を読み取り、seed値に設定
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {

		// crypt/randからReadできなかった場合の代替手段
		s = time.Now().UnixNano()
	}
	rand.Seed(s)

	// 0 ~ 99までの乱数
	n = rand.Intn(100)
	fmt.Println(n)
}
