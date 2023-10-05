package main

import "fmt"

func main() {
	inta := 5
	adinta := &inta

	bdinta := &inta
	*bdinta = 9

	fmt.Println("*bdintaを変更:")
	fmt.Printf("inta の値は %d です\n", inta)
	fmt.Printf("abinta のアドレスは %d です\n", *adinta)

	fmt.Printf("adinta のアドレスは %p です\n", &adinta)
	fmt.Printf("bdinta のアドレスは %p です\n", &bdinta)
}
