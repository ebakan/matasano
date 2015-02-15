package main

import (
	"./utils"
	"encoding/hex"
	"fmt"
)

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	fmt.Println("Input:", input)
	target, _ := hex.DecodeString(input)
	utils.BruteForceXor(target)
	fmt.Println("Output:", string(target))
}
