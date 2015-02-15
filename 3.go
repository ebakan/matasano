package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
)

var EnglishFrequencies = [256]float32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.0812, 0.0149, 0.0271, 0.0432, 0.1202, 0.023, 0.0203, 0.0592, 0.0731, 0.001, 0.0069, 0.0398, 0.026099999999999998, 0.0695, 0.0768, 0.0182, 0.0011, 0.0602, 0.06280000000000001, 0.091, 0.0288, 0.0111, 0.0209, 0.0017000000000000001, 0.021099999999999997, 0.0007000000000000001, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

func SingleXor(input []byte, cipher byte) []byte {
	byteLength := len(input)
	xor := make([]byte, byteLength)
	for i := 0; i < byteLength; i++ {
		xor[i] = input[i] ^ cipher
	}
	return xor
}

func ByteFrequency(input []byte) [256]float32 {
	var buckets [256]float32
	incr := 1 / float32(len(input))
	for i := 0; i < len(input); i++ {
		buckets[input[i]] += incr
	}
	return buckets
}

func DistanceToEnglish(input []byte) float32 {
	lowercase := bytes.ToLower(input)
	frequencies := ByteFrequency(lowercase)
	var sum float64 = 0
	for i := 0; i < 256; i++ {
		sum += math.Pow(float64(frequencies[i]-EnglishFrequencies[i]), float64(2))
	}
	return float32(math.Sqrt(sum / 256))
}

func SingleByteXor(input []byte, cipher byte) {
	for i := 0; i < len(input); i++ {
		input[i] ^= cipher
	}
}

func BruteForceXor(input []byte) {
	var bestScore float32 = 99999
	var bestCipher byte
	for i := byte(0); i < 255; i++ {
		SingleByteXor(input, i)
		score := DistanceToEnglish(input)
		if score < bestScore {
			bestScore = score
			bestCipher = i
		}
		SingleByteXor(input, i)
	}
	SingleByteXor(input, bestCipher)
}

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	fmt.Println("Input:", input)
	target, _ := hex.DecodeString(input)
	BruteForceXor(target)
	fmt.Println("Output:", string(target))
}
