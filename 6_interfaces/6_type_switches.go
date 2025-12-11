package main

import "fmt"

func findingType(data interface{}) {

	switch data.(type) {
	case int:
		fmt.Println("Integer:", data)
	case float64:
		fmt.Println("Float64:", data)
	case string:
		fmt.Println("String:", data)
	}
}

func main() {

	findingType(1)
	findingType(12.33)
	findingType("Avi")

}
