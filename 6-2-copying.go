package main

import "fmt"

func main() {
	inta := 5
	intb := inta
	inta = 7
	fmt.Printf("inta は %d です\n", inta)
	fmt.Printf("intb は %d です\n", intb)

	intb = 9
	fmt.Printf("inta は %d です\n", inta)
	fmt.Printf("intb は %d です\n", intb)

	fmt.Printf("inta のアドレスは %p です\n", &inta)
	fmt.Printf("intb のアドレスは %p です\n", &intb)
}
