package main

import (
	"fmt"
	"time"
)

func greet(text string) {
	fmt.Println(text)
}

func slowGreet(text string) {
	time.Sleep(3 * time.Second)
	fmt.Println(text)
}

func main() {

	go greet("Hii...")
	go greet("How are you..")
	go slowGreet("Nice to meet you")
	go greet("i am fine..")
}
