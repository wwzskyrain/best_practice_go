package main

type book struct {
	author, title string
}

type maganize struct {
	title string
	issue int
}

type printor interface {
	print()
}

func (b book) print() {
	println("book.title =", b.title)
}

func (m maganize) print() {
	println("magnize.title =", m.title)
}

func TestPolymorphism() {
	b := book{title: "daliry of yueyi", author: "yueyi"}
	m := maganize{title: "odd sunday?", issue: 1}

	b.print()
	m.print()

	// p指向了book或者magazine，对同样的指令print()表现出不同的行为，这就是多态‘polymorphism’
	var p printor
	p = &b
	p.print()
	p = &m
	p.print()
}
