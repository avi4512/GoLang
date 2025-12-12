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

	branch.output()

}
