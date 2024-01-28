package main

import "fmt"

func allocSlice1(min, high int) []int {
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("slice b: len(%d), cap(%d), elements(%v)\n",
		len(b), cap(b), b)
	return b[min:high]
}

func testAllocSlice1() {
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("slice b: len(%d), cap(%d), elements(%v)\n",
		len(b), cap(b), b)
	b1 := b[3:7]
	fmt.Printf("slice b1: len(%d), cap(%d), elements(%v)\n",
		len(b1), cap(b1), b1)
	fmt.Printf("slice b1: len(%d), cap(%d), elements(%v)\n",
		len(b1), cap(b1), b1)
	b2 := b1[:6]
	fmt.Printf("slice b2: len(%d), cap(%d), elements(%v)\n",
		len(b2), cap(b2), b2)
}

func allocSlice2(min, high int) []int {
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("slice b: len(%d), cap(%d), elements(%v)\n",
		len(b), cap(b), b)

	nb := make([]int, high-min) //cap和len都是high-min
	//通过声明新slice和copy，来完成复制
	copy(nb, b[min:high])
	return nb
}

func testAllocSlice2() {
	b1 := allocSlice2(3, 7)
	fmt.Printf("slice b1: len(%d), cap(%d), elements(%v)\n",
		len(b1), cap(b1), b1)

	// 这里会越界而panic
	b2 := b1[:6] //这里的b1的len和cap都是4=7-3了。
	fmt.Printf("slice b2: len(%d), cap(%d), elements(%v)\n",
		len(b2), cap(b2), b2) // panic: runtime error: slice bounds out of range [:6] with capacity 4
}

func main() {
	//testSliceOpt()
	//testAllocSlice1()
	//testAllocSl ice2()
	//towVeryInterestingFeatureAboutSlice()
	testAppend()
}

func testSliceOpt() {
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice("b", b)
	var b1 = b[0:]
	printSlice("b1", b1)
	var b2 = b[4:]
	printSlice("b2", b2)
	var b3 = b[4:6] //b3的cap将是6，你敢信吗
	printSlice("b3", b3)
}

func printSlice(name string, b []int) {
	fmt.Printf("slice [%s]: len(%d), cap(%d), elements(%v)\n", name,
		len(b), cap(b), b)
}

// 弄明白了，原来slice有这两个特性呢
func towVeryInterestingFeatureAboutSlice() {
	var a = make([]int, 4, 8)
	//可以初始化一个len为4，而cap为8的slice。这4个元素和剩下的4个空cap有什么区别呢？不能直接访问吗，对的
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[3])
	//fmt.Println(a[4]) //会越界而panic
	//fmt.Println(a[5]) //会越界而panic
	//fmt.Println(a[6]) //会越界而panic
	//fmt.Println(a[7]) //会越界而panic
	printSlice("a", a)
	var b = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	printSlice("b", b)
	var c = b[2:9] // 哦哦哦，原来如此，c的len=9-7，而cap=2*len
	printSlice("c", c)
	//fmt.Println(c[0])
	//fmt.Println(c[1])
	//fmt.Println(c[2])
	//fmt.Println(c[3])
	//fmt.Println(c[4]) //会越界panic
	var d = c[:9] //因为9小于cap(c)所以，所以不会越界
	printSlice("d", d)

	//总结，访问时，len是界限；reslice时，cap是界限.

}

func testAppend() {
	var s = make([]int, 0)
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)
	s = append(s, 5) // len=5， cap=8，
	printSlice("s", s)
	// 对于slice，可以用append函数来追加，而无需担心越界。
	// 那么，队列如何实现呢？
}
