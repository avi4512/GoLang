package main

import "fmt"

func main() {

	arr := []int{1, 2, 3}

	doubleRes := transFormer(arr, func(num int) int { return num * 2 })
	fmt.Println("Double res using Anonumus Funnction:", doubleRes)

	tripleRes := transFormer(arr, func(num int) int { return num * 3 })
	fmt.Println("Triple res using Anonumus Funnction:", tripleRes)

}

func transFormer(ar []int, something func(int) int) []int {

	dNums := []int{}
	for _, val := range ar {

		dNums = append(dNums, something(val))

	}
	return dNums

}
