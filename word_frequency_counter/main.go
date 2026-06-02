package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type dict struct {
	word string
	freq int
}

func sorter(wordsFreq map[string]int) []dict {

	var freqTuple []dict
	for w, f := range wordsFreq {
		freqTuple = append(freqTuple, dict{w, f})
	}

	sort.Slice(freqTuple, func(i, j int) bool {
		return freqTuple[i].freq > freqTuple[j].freq
	})
	return freqTuple
}

func main() {
	filesPath := "./txtFiles"
	files, err := os.ReadDir(filesPath)

	if err != nil {
		fmt.Print("No Files found inside this directory!")
		return
	}

	txtFiles := []string{}

	for _, entry := range files {
		txtFiles = append(txtFiles, entry.Name())
	}

	wordsFrequency := map[string]int{}
	wordsFileName := map[string]map[string]int{}

	for _, fileName := range txtFiles {
		fullPath := filepath.Join(filesPath, fileName)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		words := strings.Fields(string(content))

		for _, word := range words {
			wordsFrequency[word] += 1

			if wordsFileName[word] == nil {
				wordsFileName[word] = make(map[string]int) 
			}
			wordsFileName[word][fileName] += 1
		}

	}

	globalSortedDict := sorter(wordsFrequency)

	for _, wordFreqTuple := range globalSortedDict[:10] {
		maxFile := ""
		maxCount := 0

		for fileName, freq := range wordsFileName[wordFreqTuple.word] {
			if freq > maxCount {
				maxFile = fileName
				maxCount = freq
			}
		}
		fmt.Printf("Word: %v	; File: %v	; Frequency: %v \n", wordFreqTuple.word, maxFile, maxCount)
	}
}
