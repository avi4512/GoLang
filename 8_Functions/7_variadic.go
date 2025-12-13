package main

import "fmt"

func main() {

	res := sumNums("Variadic", 20, 30, 40)
	fmt.Println("Sum of numbers:", res)

}

func sumNums(firstNum string, numbers ...int) int {

	sum := 0
	for _, val := range numbers {
		sum = sum + val
	}
	fmt.Println("first Val:", firstNum)
	return sum

}
