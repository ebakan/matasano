package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	fmt.Println("Input 1:", input1)
	fmt.Println("Input 2:", input2)
	bytes1, _ := hex.DecodeString(input1)
	bytes2, _ := hex.DecodeString(input2)
	byteLength := len(bytes1)
	xor := make([]byte, byteLength)
	for i := 0; i < byteLength; i++ {
		xor[i] = bytes1[i] ^ bytes2[i]
	}
	b64 := hex.EncodeToString(xor)
	fmt.Println("Output:", b64)
}
