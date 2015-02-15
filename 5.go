package main

import (
	"encoding/hex"
	"fmt"
)

func RepeatingKeyXor(input []byte, key []byte) []byte {
	length := len(input)
	keyLength := len(key)
	output := make([]byte, length)
	for i := 0; i < length; i++ {
		output[i] = input[i] ^ key[i%keyLength]
	}
	return output
}

func main() {
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	encrypted := RepeatingKeyXor([]byte(plaintext), []byte(key))
	fmt.Printf("Input:\n%s\n", plaintext)
	fmt.Printf("Output:\n%s\n", hex.EncodeToString(encrypted))
}
