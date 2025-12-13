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
		dnum = append(dnum, val*2)
	}
	return dnum
}
