package main

import "fmt"

func generics(x, y interface{}) interface{} {

	xint, isXint := x.(int)
	yint, isYint := y.(int)

	if isXint && isYint {
		return xint + yint
	}

	xfloat, isXfloat := x.(float64)
	yfloat, isYfloat := y.(float64)

	if isXfloat && isYfloat {
		return xfloat + yfloat
	}

	xstring, isXstring := x.(string)
	ystring, isYstring := y.(string)

	if isXstring && isYstring {
		return xstring + ystring
	}
	return nil

}

func main() {

	res := generics("Good ", "Morning")
	fmt.Println(res)

}
