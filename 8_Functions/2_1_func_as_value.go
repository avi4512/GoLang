package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4}
	double := doubleNumbers(&numbers)
	fmt.Println("Orginal:", numbers)
	fmt.Println("Doubled:", double)

}

func doubleNumbers(arr *[]int) []int {

	dnum := []int{}
	for _, val := range *arr {
		dnum = append(dnum, doubleNum(val)) // here we are passing function as value
	}
	return dnum
}

func doubleNum(num int) int {
	return num * 2
}

// func tripleNum(num int) int {
// 	return num * 3
// }
