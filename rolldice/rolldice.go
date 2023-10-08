package main

import (
	"fmt"
	"math/rand"
	"time"
)

type dice struct {
	val int
}

func rollDice(d dice, c chan int) {
	rand.Seed(time.Now().UnixNano())
	for {
		time.Sleep(10 * time.Millisecond)
		v := rand.Intn(10)
		if d.val == v {
			fmt.Printf("出たー！\n")
			break
		} else {
			fmt.Printf("%dか……%dではないな\n", v, d.val)
		}
	}
	c <- d.val
}

func main() {
	d1 := dice{2}
	d2 := dice{6}

	c := make(chan int)

	go rollDice(d1, c)
	go rollDice(d2, c)

	x, y := <-c, <-c

	fmt.Printf("%dが出ました\n", x)
	fmt.Printf("%dが出ました\n", y)

}
