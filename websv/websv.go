package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.HandlerFunc(makePage))
	http.ListenAndServe(":8782", nil)
}

func makePage(resw http.ResponseWriter, req *http.Request) {
	counter, err := ioutil.ReadFile("indexhtml.txt")
	if err != nil {
		fmt.Println("テンプレートファイルの読み込みエラーです", err)
		os.Exit(1)
	} else {
		templatestr := string(counter)
		temple := template.Must(template.New("sendname").Parse(templatestr))
		temple.Execute(resw, req.FormValue("yourname"))
	}
}
