package main

import "fmt"

type input interface {
	read() string
}

type keyboard struct{}

type file struct{}

type mic struct{}

type sensor struct{}

func (k keyboard) read() string {
	return "何かキーを押してください"
}

func (f file) read() string {
	return "ファイルパスを指定してください"
}

func (f mic) read() string {
	return "ボタンを押して話してください"
}

func (s sensor) read() string {
	return "計測を開始してください"
}

func main() {

	methods := []input{
		keyboard{},
		file{},
		mic{},
		sensor{},
	}

	for i := 0; i < len(methods); i++ {
		fmt.Println(methods[i].read())
	}
}
