package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type dict struct {
	word string
	freq int
}

func main() {
	content, err := os.ReadFile("pg24681.txt")

	if err != nil {
		fmt.Print("ERROR", err)
		return
	}

	words := strings.Fields(string(content))
	frequency := map[string]int{}

	for i := range words {
		frequency[words[i]] = frequency[words[i]] + 1
	}

	var freqs []dict
	for w, f := range frequency {
		freqs = append(freqs, dict{w, f})
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].freq > freqs[j].freq
	})
	fmt.Println(freqs[:10])
}
