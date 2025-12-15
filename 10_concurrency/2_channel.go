package main

import (
	"fmt"
	"time"
)

func greet(text string) {
	fmt.Println(text)
}

func slowGreet(text string, done_bool chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(text)
	done_bool <- true
}

func main() {

	go greet("Hii...")
	go greet("How are you..")
	done_bool := make(chan bool)
	go slowGreet("Nice to meet you", done_bool)
	go greet("i am fine..")
	<-done_bool
}
