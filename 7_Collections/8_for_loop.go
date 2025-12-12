package main

import "fmt"

// aliases
type intMap map[int]string

func (m intMap) output() {
	fmt.Println(m)
}

func main() {

	branch := intMap{}
	branch[101] = "CSE"
	branch[102] = "ECE"
	branch[103] = "EEE"

	names := []string{"Avi", "Vamsi", "Malli", "Mukesh"}
	fmt.Println("---------arry's for loop--------------")
	for index, value := range names {
		fmt.Println("Index:", index, " ", "Values:", value)
	}

	fmt.Println("--------Map's For Loop ---------------")
	for key, val := range branch {
		fmt.Println("Key:", key, " ", "value:", val)
	}

}
