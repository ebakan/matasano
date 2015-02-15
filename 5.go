package main

import (
	"./utils"
	"encoding/hex"
	"fmt"
)

func main() {
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	encrypted := utils.RepeatingKeyXor([]byte(plaintext), []byte(key))
	fmt.Printf("Input:\n%s\n", plaintext)
	fmt.Printf("Output:\n%s\n", hex.EncodeToString(encrypted))
}
