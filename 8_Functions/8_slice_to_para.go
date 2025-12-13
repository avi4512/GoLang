package main

import "fmt"

func main() {

	arr := []int{10, 20, 30, 40}

	other := sumNums(1, arr...) // ... indicates splitting values
	fmt.Println(other)

}

func sumNums(firstVal int, numbers ...int) int {

	sum := 0
	for _, val := range numbers {
		sum = sum + val
		fmt.Println(val)
	}

	return sum

}
