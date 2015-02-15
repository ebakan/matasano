package main

import (
	"./utils"
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	filename := "4.txt"
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var bestScore float64 = 99999
	var bestInput []byte
	var bestCipher byte
	for scanner.Scan() {
		str := scanner.Text()
		bytes, _ := hex.DecodeString(str)
		cipher, score := utils.BruteForceXor(bytes)
		if score < bestScore {
			bestScore = score
			bestInput = bytes
			bestCipher = cipher
		}
	}
	fmt.Println("Output:", string(bestInput))
	utils.SingleByteXor(bestInput, bestCipher)
	fmt.Println("Input:", hex.EncodeToString(bestInput))
	fmt.Println("Cipher:", bestCipher)
	fmt.Println("Score:", bestScore)
}
