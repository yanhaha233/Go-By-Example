package main

import "fmt"

func main11() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	//这里我们使用 range 来对 slice 中的元素求和。 数组也可以用这种方法初始化并赋初值。
	//此处未用到索引，所以使用空白标识符_将其忽略
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
		// 此处使用到索引
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
		// range在map中迭代键值对
	}
	for k := range kvs {
		fmt.Println("key:", k)
		// range也可以只遍历map的键
	}
	for i, c := range "go" {
		fmt.Println(i, c)
		// range在字符串中迭代unicode码点。
		// 第一个返回值是字符的起始字节位置,然后第二个是字符本身。
	}
}
