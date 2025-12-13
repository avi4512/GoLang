package main

import "fmt"

func main() {

	res := factorial(6)
	fmt.Println(res)

}

func factorial(num int) int {
	if num == 1 {

		return 1
	}
	return num * factorial(num-1)
}
