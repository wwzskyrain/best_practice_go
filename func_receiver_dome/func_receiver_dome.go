package main

import "fmt"

type notifier interface {
	notify()
}
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	// 方法内，可用u直接访问，而不用再搞*u了。
	fmt.Printf("send user email to %s<%s>\n",
		u.name, u.email)
}

func sendNotify(u notifier) {
	u.notify()
}

func TestCallSendNotiry() {
	u := user{name: "yueyi", email: "1191113214@qq.com"}
	// 传u对象的话，会编译报错：
	// cannot use u (variable of type user) as notifier value in argument to sendNotify: missing method notify (notify has pointer receiver) compilerInvalidIfaceAssign
	// 就是说，不能能用u变量以notifier的value的身份来传参给sendNotify：因为notify是一个指针接受者【pointer receiver】
	// 要明确的给一个地址值。因为**编译器不一定总能拿到一个值的地址**
	// 哎，就是为了写这句话来的，这话还有一个意思，就是如果参数是值类型，那么传过来一个值或者指针都行
	// sendNotify(u)
	sendNotify(&u)

	// 当然，可以直接用u对象调用notify()方法
	u.notify()
	// 甚至也可以用u的引用来调用notify()，因为在go里面变量和变量的地址在调用上是等价的
	(&u).notify()

	// 下面我们试着剥离一下sentNotify()的函数这一层，
	// 深入到方法接受者的本质。
	// 这种声明变量在赋值的操作会触发编译告警，但是为了说明问题，这里我们就不优化了
	var n notifier
	// cannot use u (variable of type user) as notifier value in assignment: missing method notify (notify has pointer receiver)compilerInvalidIfaceAssign
	// 看到了吧，和sentNotify()中一样的报错信息，说明即使不用函数传参的形式，
	// 也一样出现接受者类型不匹配的问题。
	// 当我们面向接口编程的时候，这个问题就比较容易遇到了
	// n = u;
	n = &u
	n.notify()
}

func main() {
	TestCallSendNotiry()
}
