package main

import "fmt"

func counter() func() {
	ctr := 0
	return func() {
		ctr++
		fmt.Println(ctr)
	}
}

func main() {
	countfunc := counter()
	countfunc()
	countfunc()
	countfunc()
}
