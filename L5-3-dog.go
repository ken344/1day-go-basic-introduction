package main

import "fmt"

type dog struct {
	name   string
	group  string
	height int
}

func biggest(dogs []dog) dog {
	biggest := dogs[0]

	for i := 1; i < len(dogs); i++ {
		if dogs[i].height > biggest.height {
			biggest = dogs[i]
		}
	}
	return biggest
}

func main() {
	pome := dog{"ポメ", "ポメラニアン", 25}
	maru := dog{"マル", "マルチーズ", 22}
	shiba := dog{"シバ", "柴犬", 40}

	dogs := []dog{pome, maru, shiba}

	fmt.Printf("%sさんが最大です\n", biggest(dogs).name)
}
