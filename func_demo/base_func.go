package main

import "fmt"

type Duck struct {
	legs int
}

// 带有接受者的方法，这是入参还是出参呢？这已经不是我们熟知的方法了
// 接受者既不是入参，也不是出参，而是接受者本身。
// 如何理解接受者呢？可以把它看成是调用者，函数的属主，这样就和java的类的方法联系起来了
// 所以这里不叫函数，叫方法。
// 其实这里还可以把接受者看做入参的，及时是调用者，也是入参之一，在java中我们定义过不少这样的“外部方法”
// 那么接受者有值接受者和指针接受者的区别，我们把它当做入参，如此一来我们就能联想到在c语言编程时的“入参的值类型和指针类型”了
// so，关于函数的操作是否作用到“原本”的结论如下：
// 1.	值接受者，是入参的一个副本，对副本的任何变动都不影响“原本”。
// 2.	指针接受者，是可以通过指针访问到“原本”的，所以在函数中对指针的属性的操作其实就在操作“原本”的属性
// 还有一个问题未解决，那就是接受者的调用属性，
// 具体问题是：允许使用怎样的属主来调用方法呢（这个时候称之为方法比称之为函数更合适）
// todo
func (dk Duck) SetLegs(n int) {
	dk.legs = n
	println("方法SetLegs中，dk.legs", dk.legs)
}

func (dk *Duck) ChangeLegs(n int) {
	(*dk).legs = n
	println("方法SetLegs中，dk.legs", (*dk).legs)
}

func ShowMeTheLegs() {
	first := Duck{}
	fmt.Println("1. First duck legs: ", first.legs)
	first.SetLegs(2)
	fmt.Println("2. First duck legs: ", first.legs)
	first.ChangeLegs(4)
	fmt.Println("3. First duck legs: ", first.legs)

	second := new(Duck)
	fmt.Println("4. Second duck legs: ", second.legs)
	second.SetLegs(4)
	fmt.Println("5. Second duck legs: ", second.legs)
	second.ChangeLegs(4)
	fmt.Println("6. Second duck legs: ", second.legs)
}

func main() {
	ShowMeTheLegs()
}
