package main

import "fmt"

type vehicle struct {
	brand string
	speed int
}

type car struct {
	model string
	vehicle
}

func main() {

	c := car{
		model: "XUV",
		vehicle: vehicle{
			brand: "THAR",
			speed: 120,
		},
	}
	c.output()

}

func (c car) output() {
	fmt.Println(c.brand)
	fmt.Println(c.model)
	fmt.Println(c.speed)
}
