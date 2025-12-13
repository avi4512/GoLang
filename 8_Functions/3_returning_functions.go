package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4}
	other_numbers := []int{5, 10, 9}
	double := changeNumbers(&numbers, doubleNum)
	triple := changeNumbers(&numbers, tripleNum)
	fmt.Println("Orginal:", numbers)
	fmt.Println("Doubled:", double)
	fmt.Println("Triple:", triple)

	conditionfn1 := getConditionFunc(numbers)
	conditionfn2 := getConditionFunc(other_numbers)

	res1 := changeNumbers(&numbers, conditionfn1)
	res2 := changeNumbers(&other_numbers, conditionfn2)

	fmt.Println("conditionfn1:", res1)
	fmt.Println("conditionfn2:", res2)

}

func changeNumbers(arr *[]int, changeTo func(int) int) []int {

	dnum := []int{}
	for _, val := range *arr {
		dnum = append(dnum, changeTo(val)) // here we are passing function as value
	}
	return dnum
}

func getConditionFunc(num []int) func(int) int {

	if num[0] == 1 {
		return doubleNum
	} else {
		return tripleNum
	}
}

func doubleNum(num int) int {
	return num * 2
}

func tripleNum(num int) int {
	return num * 3
}
