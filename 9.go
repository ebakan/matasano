package main

import (
	"./utils"
	"fmt"
)

func main() {
	input := "YELLOW SUBMARINE"
	output := utils.PKCS7Padding([]byte(input), 20)
	fmt.Println("Input:", input)
	fmt.Println("Input: ", []byte(input))
	fmt.Println("Output:", output)
}
