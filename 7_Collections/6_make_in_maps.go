package main

import "fmt"

func main() {

	branch := make(map[int]string, 3)

	branch[101] = "CSE"
	branch[102] = "ECE"
	branch[103] = "EEE"

	fmt.Println(branch)

}
