package main

import "fmt"

func main() {

	id := make([]int, 3, 5)

	id = append(id, 4)
	id = append(id, 5)
	id = append(id, 6)

	fmt.Println(id)

}
