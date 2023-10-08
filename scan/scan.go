package main

import "fmt"

func main() {
	var input string
	for {
		fmt.Print("入力してください。'q'で終了:")
		fmt.Scanln(&input)

		if input == "q" {
			break
		}
		fmt.Println(input, "と入力されました")
	}
	fmt.Println("終了")
}
