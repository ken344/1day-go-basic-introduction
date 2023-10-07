package main

import (
	"fmt"
)

func counter() func() {
	ctr := 0
	fmt.Println("カウンタを初期化しました")
	fmt.Printf("カウンタのアドレスは%p\n", &ctr)
	return func() {
		ctr++
		fmt.Printf("現在のカウンタは%dです。", ctr)
		fmt.Printf("カウンタのアドレスは%p\n", &ctr)
	}
}

func main() {
	countfnc := counter()
	countfnc()
	countfnc()
	countfnc()
}
