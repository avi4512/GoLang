package main

import "fmt"

func typeVal(val interface{}) {

	_, intOk := val.(int)
	_, floatOk := val.(float64)
	_, stringOK := val.(string)

	if intOk {
		fmt.Println("int")
	} else if floatOk {
		fmt.Println("float")
	} else if stringOK {
		fmt.Println("string")
	} else {
		fmt.Println("Invalid type!!")
	}

}

func main() {

	typeVal(1)
	typeVal(1.2)
	typeVal("Hii.")

}
