package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	key := "YELLOW SUBMARINE"
	filename := "7.txt"
	data, _ := ioutil.ReadFile(filename)
	bytes, _ := base64.StdEncoding.DecodeString(string(data))
	block, _ := aes.NewCipher([]byte(key))
	blockSize := block.BlockSize()
	consumableBytes := bytes
	for len(consumableBytes) > 0 {
		block.Decrypt(consumableBytes, consumableBytes)
		consumableBytes = consumableBytes[blockSize:]
	}
	fmt.Println("Decrypted:")
	fmt.Println(string(bytes))
}
