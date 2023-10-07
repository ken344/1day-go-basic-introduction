package main

import "fmt"

type guest struct {
	name string
	age  int
	job  string
}

func (g *guest) present() string {
	return fmt.Sprintf("%sさんは%d歳で%sです", g.name, g.age, g.job)
}

func main() {
	legs := map[string]int{
		"bird":        2,
		"cat":         4,
		"grasshopper": 6,
		"octopus":     8,
		"squid":       10,
	}
	for k, v := range legs {
		fmt.Printf("A %s has %d legs\n", k, v)
	}

	rooms := make(map[int]guest)
	rooms[101] = guest{"山田", 29, "会社員"}
	rooms[208] = guest{"川崎", 21, "学生"}
	rooms[312] = guest{"高橋", 45, "公務員"}

	for k, v := range rooms {
		fmt.Printf("%d号室:", k)
		fmt.Println(v.present())
	}
}
