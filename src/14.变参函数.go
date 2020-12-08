package main

import "fmt"

func sum(nums ...int) { // 这个函数接受任意数量的 int 作为参数
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main14() {
	sum(1, 2) // 变参函数使用常规的调用方式，传入独立的参数
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// 如果你有一个含有多个值的 slice，想把它们作为参数使用， 你需要这样调用 func(slice...)。
	sum(nums...)
}
