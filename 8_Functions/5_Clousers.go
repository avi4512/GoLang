package main

import "fmt"

func main() {

	arr := []int{1, 2, 3}

	double := changeTransFromer(2)          //using clouser function
	double_res := transFormer(&arr, double) // passing the func as argument

	triple := changeTransFromer(3)
	triple_res := transFormer(&arr, triple)

	fmt.Println("Changing using Clousers: for Double:", double_res)
	fmt.Println("Changing using Clousers: for triple:", triple_res)

}

//clouser(retuning anonumus function)
// creating clouser
func changeTransFromer(factor int) func(int) int {
	return func(val int) int {
		return factor * val
	}
}

func transFormer(ar *[]int, something func(int) int) []int {

	dNums := []int{}
	for _, val := range *ar {

		dNums = append(dNums, something(val))

	}
	return dNums
}
