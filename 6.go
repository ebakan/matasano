package main

import (
	"./utils"
	"container/heap"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

type Pair struct {
	key   int
	value float32
}

type PairHeap []Pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HammingDistance(arr1 []byte, arr2 []byte) int {
	sum := 0
	for i := 0; i < len(arr1); i++ {
		xor := arr1[i] ^ arr2[i]
		for xor != 0 {
			sum += int(xor & 1)
			xor = xor >> 1
		}
	}
	return sum
}

func main() {
	filename := "6.txt"
	data, _ := ioutil.ReadFile(filename)
	bytes, _ := base64.StdEncoding.DecodeString(string(data))
	numBytes := len(bytes)
	distanceHeap := &PairHeap{}
	heap.Init(distanceHeap)
	for keysize := 2; keysize <= 40; keysize++ {
		slice1 := bytes[:keysize]
		slice2 := bytes[keysize : keysize*2]
		distance := float32(HammingDistance(slice1, slice2)) / float32(keysize)
		pair := Pair{key: keysize, value: distance}
		heap.Push(distanceHeap, pair)
	}
	var bestCipher []byte
	var bestDecrypted []byte
	bestScore := 999.0
	for trial := 0; trial < 19; trial++ {
		pair := heap.Pop(distanceHeap).(Pair)
		keysize := pair.key
		blocksT := make([][]byte, keysize)
		for i := 0; i < keysize; i++ {
			blocksize := (numBytes-i-1)/keysize + 1
			blocksT[i] = make([]byte, blocksize)
			for j := 0; j < blocksize; j++ {
				blocksT[i][j] = bytes[j*keysize+i]
			}
		}
		cipher := make([]byte, keysize)
		for i := 0; i < keysize; i++ {
			cipher[i], _ = utils.BruteForceXor(blocksT[i])
		}
		decrypted := utils.RepeatingKeyXor(bytes, cipher)
		score := utils.DistanceToEnglish(decrypted)
		if score < bestScore {
			bestCipher = cipher
			bestDecrypted = decrypted
			bestScore = score
		}
	}
	fmt.Println("Cipher:")
	fmt.Println(string(bestCipher))
	fmt.Println("Decrypted:")
	fmt.Println(string(bestDecrypted))
}
