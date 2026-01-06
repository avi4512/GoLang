package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFromFile() {

	byte_data, err := os.ReadFile("amount_file.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	content := strings.TrimSpace(string(byte_data))

	arr_vals := strings.Split(content, "\n")

	for _, val := range arr_vals {

		float_val, err := strconv.ParseFloat(val, 64)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(float_val)
	}

}

func main() {
	readFromFile()
}
