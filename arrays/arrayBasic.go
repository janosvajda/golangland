package main

import "fmt"

func main() {

	//It defines an nums of integers.
	nums := []int{0, 11, 22, 33, 44, 55, 555}

	fmt.Println("The whole array: ", nums)

	//It iterates over key/pair values.
	for i, num := range nums {
		fmt.Println("Index: ", i, "Value: ", num)
	}

}
