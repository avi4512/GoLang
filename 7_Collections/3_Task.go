package main

import "fmt"

func main() {

	//1. Creating array with 3 Hobbies and printing
	hobbies := [3]string{"Reading", "Writting", "Learning"}
	fmt.Println("1.Hobbies:", hobbies)

	//2.1 first element
	fmt.Println("2.First Element:", hobbies[0])

	//2.2 combine of 2 and 3 elements
	combine_list := hobbies[1:3]
	fmt.Println("2.Combine of 2 and 3 elements:", combine_list)

	//3 slice based on first element
	slice_first := hobbies[:2]
	//slice_first := hobbies[0:2]
	fmt.Println("3.Slicing 1 and 2 elements:", slice_first)

	//4. Re-slice based on (task-3) to get 2 and 3 element
	slice_first = slice_first[1:3]
	fmt.Println("4. Re-slice based on (task-3) to get 2 and 3 element:", slice_first)

	//5.course goal(atleat 2)
	dynamic_array := []string{"Complete Go", "Read one Page.."}
	fmt.Println("Dynamic Array with 2 Goal's", dynamic_array)

	//6.adding 3rd goal to the dynamic array
	dynamic_array = append(dynamic_array, "Bhagavath Geetha")
	fmt.Println("Adding 3rd goal to the dynamic Array:", dynamic_array)

	//7.creating struct and adding 2 product's

	shop := []Products{
		{101, "Soap", 45.67},
		{102, "Paste", 67.89},
	}
	fmt.Println("Products using dynamic Struct's:", shop)

	//7.2 adding one product existing list
	newProduct := Products{103, "Shampoo", 2.45}
	shop = append(shop, newProduct)
	fmt.Println("Adding new product:", shop)

}

type Products struct {
	id    int
	name  string
	price float64
}
