package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
	// fact 函数在到达 fact(0) 前一直调用自身。
}

func main16() {
	fmt.Println(fact(7))
}
