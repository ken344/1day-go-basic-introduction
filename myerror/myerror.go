package main

import "fmt"

type myerror struct {
	msg string
}

func (e *myerror) Error() string {
	return e.msg
}

func read(word string) (string, error) {
	if len(word) > 0 {
		return word + "...と言いましたね", nil
	}
	return "", &myerror{"読み取れませんでした"}
}

func main() {
	words := []string{"Hey", "", "ゲ"}

	for i := 0; i < len(words); i++ {
		result, err := read(words[i])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(result)
		}
	}
}
