package main

import "fmt"

func main() {

	nums := []int{0, 11, 22, 33, 44, 55, 555}

	//It defines an nums of integers.
	fmt.Println("The whole array: ", nums)

	//It iterates over key/pair values.
	for i, num := range nums {
		fmt.Println("Index: ", i, "Value: ", num)
	}

}
