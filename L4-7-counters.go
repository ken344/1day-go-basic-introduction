package main

import "fmt"

func counter1() func() {
	ctr := 0
	return func() {
		ctr++
		fmt.Println(ctr)
	}
}

func main() {
	countfunc := counter1()
	countfunc()
	countfunc()
	countfunc()
}
