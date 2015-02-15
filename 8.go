package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	filename := "8.txt"
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		bytes, _ := hex.DecodeString(str)
		m := make(map[string]bool)
		for len(bytes) > 0 {
			chunk := bytes[:16]
			bytes = bytes[16:]
			if _, ok := m[string(chunk)]; ok {
				fmt.Println("Found ECB text:")
				fmt.Println(str)
				break
			} else {
				m[string(chunk)] = true
			}
		}
	}
}
