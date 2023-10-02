package main

import "fmt"

type dog struct {
	name   string
	group  string
	height int
}

func (d dog) describe() string {
	dscr := "わたしは" + d.group
	dscr += "、名は" + d.name
	return dscr
}

func (d dog) bigger(d2 dog) string {

	dscr := d.describe()
	dscr += "、私は" + d2.name

	diff := d.height - d2.height
	switch {
	case diff > 5:
		dscr += "より大きい"
	case diff < -5:
		dscr += "より小さい"
	default:
		dscr += "と同じくらいである"
	}
	return dscr
}
func main() {
	maru := dog{"マル", "マルチーズ", 22}
	shiba := dog{"シバ", "柴犬", 40}

	fmt.Println(maru.bigger(shiba))

}
