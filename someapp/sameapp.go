package main

import "fmt"

type someapp struct {
	username string
	use      int
	isopen   bool
}

func (app *someapp) open() {
	if app.isopen {
		fmt.Printf("%sさんのアプリはすでに開いています\n\n", app.username)
	} else {
		app.use++
		if app.use > 2 {
			fmt.Printf("%sさんの使用期間は満了です。購入してください\n\n", app.username)
		} else {
			fmt.Printf("こんにちは%sさん、%d回目のご使用になります\n", app.username, app.use)
			app.isopen = true
		}
	}
}

func (app *someapp) close() {
	if app.isopen {
		fmt.Printf("%sさんのアプリを終了します\n\n", app.username)
		app.isopen = false
	}
}

func newapp(username string) someapp {
	fmt.Printf("%sさんのアプリを開きます\n\n", username)
	return someapp{username, 1, true}
}

func main() {
	fmt.Println("[ringoさんがアプリを開きます]")
	rapp := newapp("ringo")
	rapp.close()

	fmt.Println("[ringoさんがアプリを再び開きます]")
	rapp.open()
	rapp.close()

	fmt.Println("[mikanさんがアプリを開きます]")
	mapp := newapp("mikan")
	mapp.close()

	fmt.Println("[mikanさんがアプリを再び開きます]")
	mapp.open()

	fmt.Println("[ringoさんがアプリを3回目に開きます]")
	rapp.open()

	fmt.Println("[mikanさんがアプリを3回目に開きます]")
	mapp.open()
}
