package main

import "fmt"

func main() {

	//creation of array

	prices := [4]float64{12.34, 15.43, 16.78, 23.56}
	slicePrice := prices[1:3]
	array_f := prices[:3]
	array_l := prices[2:]

	fmt.Println("Original Array:", prices)
	fmt.Println("Sliced Array:", slicePrice)
	fmt.Println("Capacity of silced Array:", cap(slicePrice))
	fmt.Println("Slice from beging to given index:", array_f)
	fmt.Println("Slice from given index to last:", array_l)
}
