package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4}
	double := changeNumbers(&numbers, doubleNum)
	triple := changeNumbers(&numbers, tripleNum)
	fmt.Println("Orginal:", numbers)
	fmt.Println("Doubled:", double)
	fmt.Println("Triple:", triple)

}

func changeNumbers(arr *[]int, changeTo func(int) int) []int {

	dnum := []int{}
	for _, val := range *arr {
		dnum = append(dnum, changeTo(val)) // here we are passing function as value
	}
	return dnum
}

func doubleNum(num int) int {
	return num * 2
}

func tripleNum(num int) int {
	return num * 3
}
