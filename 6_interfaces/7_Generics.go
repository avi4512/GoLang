package main

import "fmt"

func generics[t int | float64 | string](x, y t) t {

	return x + y

}
func output(value interface{}) {
	fmt.Println("Result: ", value)
}
func main() {

	intRes := generics(10, 20)
	floarRes := generics(1.23, 3.45)
	stringRes := generics("Hello ", "world")

	output(intRes)
	output(floarRes)
	output(stringRes)
}
