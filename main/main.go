package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

func main() {
	files := []string{} // Add your file names here
	trigramFreq := make(map[string]int)

	if len(os.Args) > 1 {
		// Read from command-line arguments
		for _, filename := range os.Args[1:] {
			files = append(files, filename)

		}
	}

	var wg sync.WaitGroup

	for _, filename := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			processFile(filename, trigramFreq)
		}(filename)
	}

	wg.Wait()

	// Convert trigram frequencies to a slice for sorting
	trigramSlice := make([]trigramEntry, 0, len(trigramFreq))
	for trigram, freq := range trigramFreq {
		trigramSlice = append(trigramSlice, trigramEntry{Trigram: trigram, Frequency: freq})
	}

	// Sort trigram entries by frequency in descending order
	sort.Slice(trigramSlice, func(i, j int) bool {
		return trigramSlice[i].Frequency > trigramSlice[j].Frequency
	})

	// Print sorted trigram frequencies
	for _, entry := range trigramSlice[:100] {
		fmt.Printf("%s: %d\n", entry.Trigram, entry.Frequency)
	}
}

type trigramEntry struct {
	Trigram   string
	Frequency int
}

func processFile(filename string, trigramFreq map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file %s: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordRegex := regexp.MustCompile(`\b[\p{L}\p{Nd}'’]+\b`)

	words := make([]string, 0, 3)

	for scanner.Scan() {
		word := wordRegex.FindString(scanner.Text())
		if word != "" {
			word = strings.ToLower(word)
			words = append(words, word)
			if len(words) == 3 {
				trigram := strings.Join(words, " ")
				trigramFreq[trigram]++
				words = words[1:]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %s: %v", filename, err)
	}
}
