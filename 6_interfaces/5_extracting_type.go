package main

import "fmt"

func extractType(data interface{}) {

	intValue, ok := data.(int)
	if ok {
		fmt.Println("Integer:", intValue)
		return
	}
	floatValue, ok := data.(float64)
	if ok {
		fmt.Println("Float64:", floatValue)
		return
	}
	stringValue, ok := data.(string)
	if ok {
		fmt.Println("String:", stringValue)
		return
	}

}

func main() {

	extractType(1)
	extractType(12.33)
	extractType("Avi")

}
