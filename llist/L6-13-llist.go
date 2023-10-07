package main

import "fmt"

type element struct {
	value string
	next  *element
}

type llist struct {
	head *element
	tail *element
}

func (list *llist) add(newvalue string) {

	elm := element{newvalue, nil}

	if list.head == nil {
		list.head = &elm
	} else {
		elm.isnextto(list.tail)
	}
	list.tail = &elm
}

func (elm *element) isnextto(prevelm *element) {
	prevelm.next = elm
}

func (elm *element) getValue() string {
	return elm.value
}

func (list *llist) listall() {
	elm := list.head
	for {
		fmt.Println(elm.getValue())
		if elm.next == nil {
			break
		}
		elm = elm.next
	}
}

func main() {
	mylist := llist{nil, nil}
	mylist.add("クマ")
	mylist.add("マンドリル")
	mylist.add("ルリビタキ")
	mylist.add("キリン")
	mylist.add("あっ")

	mylist.listall()

}
