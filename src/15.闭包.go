package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
	//intSeq 函数返回一个在其函数体内定义的匿名函数
	//返回的函数使用闭包的方式 隐藏 变量 i
	//返回的函数 隐藏 变量 i 以形成闭包
}

func main15() {
	nextInt := intSeq()

	//我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt
	//这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时，都会更新 i 的值
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下。
	newInts := intSeq()
	fmt.Println(newInts())
}
