package main

import "fmt"

type animaldata interface {
	describe() string
}

type spec struct {
	name  string
	group string
}

func (sp spec) identify() string {
	return sp.group + "の" + sp.name + "さん"
}

type dog struct {
	sp     spec
	weight int
}

type cat struct {
	sp    spec
	btype string
}

func (d dog) classify() string {
	if d.weight > 25 {
		return "大型犬"
	}
	if d.weight < 10 {
		return "小型犬"
	}
	return "中型犬"
}

func (d dog) describe() string {
	dscr := d.sp.identify()
	dscr += "は" + d.classify()
	return dscr
}

func newcat(name string, group string, typenum int) cat {
	sp := spec{name, group}

	if typenum < 0 || typenum > 5 {
		typenum = 5
	}

	btypes := [6]string{"オリエンタル", "コビー", "セミコビー", "フォーリン", "セミフォーリン", "サブスタンシャル"}

	return cat{sp, btypes[typenum]}
}

func (c cat) describe() string {
	dscr := c.sp.identify() + "、"
	dscr += "ボディタイプは" + c.btype

	return dscr
}

func main() {
	animals := []animaldata{
		dog{spec{"コゲ", "コーギー"}, 15},
		newcat("クロ", "雑種", 2),
		newcat("プリンス", "ペルシャ", 1),
		dog{spec{"ルドルフ", "ハスキー"}, 27},
	}

	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i].describe())
	}
}
