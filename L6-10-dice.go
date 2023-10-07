package main

import "fmt"

type roll struct {
	round int
	score int
}

func begin() *roll {
	fmt.Println("サイコロを投げます")
	return &roll{0, 0}
}

func throw(r *roll, x int) {
	r.round++

	if x%2 == 0 {
		r.score++
	}
	fmt.Printf("%d回目: スコア=%d\n",
		r.round, r.score)
}

func main() {
	myroll := begin()
	throw(myroll, 6)
	throw(myroll, 5)
	throw(myroll, 2)
	throw(myroll, 4)
	throw(myroll, 3)
}
