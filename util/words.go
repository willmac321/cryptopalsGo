package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetWordList() []string {
	file, _ := filepath.Abs("../data/words.txt")
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("File read error", err)
	}
	rv := strings.Split(strings.TrimSpace(string(data[:])), "\n")
	for i, e := range rv {
		rv[i] = strings.TrimSpace(e)
	}
	return rv
}

func ReadFile(fp string) []string {
	file, _ := filepath.Abs(fp)
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("File read error", err)
	}
	rv := strings.Split(strings.TrimSpace(string(data[:])), "\n")
	return rv
}

func FindByteFrequency(bytes []byte) map[byte]int {
	compareMap := make(map[byte]int)
	for i := 0; i < len(bytes); i++ {
		c := bytes[i]
		compareMap[c] = compareMap[c] + 1
		fmt.Printf("byte %v appears %d times\n", string(c), compareMap[c])
	}
	return compareMap
}

func CharWeights() map[byte]int {
	str := "ETAOIN SHRDLU"
	weights := map[byte]int{}
	for i, c := range str {
		weights[byte(c)] = len(str) - i
		lower := string(c)
		lower = strings.ToLower(lower)
		lowerB := []byte(lower)
		weights[lowerB[0]] = len(str) - i
	}
	return weights
}

func GetWeights(bytes []byte) int {
	weights := CharWeights()
	w := 0
	for _, v := range bytes {
		if val, ok := weights[v]; ok {
			w += val
		}
	}
	return w
}

func HammingDistance(bytes1 []byte, bytes2 []byte) int {
	// throw
	if len(bytes1) != len(bytes2) {
		log.Fatal("Cant find hamming dist, string lengths are not equal")
	}
	dist := 0
	for i := 0; i < len(bytes1); i++ {
		// xor to find the different bits
		val := bytes1[i] ^ bytes2[i]
		// while val has different bits
		for val > 0 {
			val &= (val - 1)
			dist++
		}
	}
	return dist
}

func HammingDistanceStr(str1 string, str2 string) int {
	return HammingDistance([]byte(str1), []byte(str2))
}
