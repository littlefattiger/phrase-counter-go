package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

func main() {
	files := []string{} // Add the file names here
	trigramFreq := make(map[string]int)
	var mu sync.Mutex

	if len(os.Args) > 1 {
		// Read from command-line arguments
		for _, filename := range os.Args[1:] {
			files = append(files, filename)

		}
		var wg sync.WaitGroup

		for _, filename := range files {
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()
				processFile(filename, trigramFreq, &mu)
			}(filename)
		}

		wg.Wait()

	} else {
		fmt.Println("Reading input from pipe:")

		inputBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}

		processContent(inputBytes, trigramFreq, &mu)
	}

	trigramSlice := make([]trigramEntry, 0, len(trigramFreq))
	for trigram, freq := range trigramFreq {
		trigramSlice = append(trigramSlice, trigramEntry{Trigram: trigram, Frequency: freq})
	}
	sortTrigramSlice(trigramSlice)
	for _, entry := range trigramSlice[:5] {
		fmt.Printf("%s: %d\n", entry.Trigram, entry.Frequency)
	}
}

type trigramEntry struct {
	Trigram   string
	Frequency int
}

func sortTrigramSlice(slice []trigramEntry) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Frequency > slice[j].Frequency
	})
}

func processFile(filename string, trigramFreq map[string]int, mu *sync.Mutex) {
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
				mu.Lock()
				trigramFreq[trigram]++
				mu.Unlock()
				words = words[1:]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %s: %v", filename, err)
	}
}

func processContent(content []byte, trigramFreq map[string]int, mu *sync.Mutex) {
	// Convert content to lowercase and split into words
	text := strings.ToLower(string(content))
	wordRegex := regexp.MustCompile(`\b[\p{L}\p{Nd}'’]+\b`)
	words := wordRegex.FindAllString(text, -1)

	// Calculate three-word sequences and their frequencies
	for i := 0; i <= len(words)-3; i++ {
		trigram := strings.Join(words[i:i+3], " ")
		mu.Lock()
		trigramFreq[trigram]++
		mu.Unlock()
	}

}
