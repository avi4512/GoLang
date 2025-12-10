package main

import "fmt"

type vehicle struct {
	brand string
	speed int
}

type car struct {
	model string
	v     vehicle
}

func main() {

	c := car{
		model: "XUV",
		v: vehicle{
			brand: "THAR",
			speed: 120,
		},
	}
	c.output()

}

func (c car) output() {
	fmt.Println(c.v.brand)
	fmt.Println(c.model)
	fmt.Println(c.v.speed)
}
