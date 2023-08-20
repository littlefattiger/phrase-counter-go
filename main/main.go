package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		// Read from command-line arguments
		for _, filename := range os.Args[1:] {
			processFile(filename)
		}
	} else {
		// Read from standard input (pipe)
		fmt.Println("Reading input from pipe:")

		// Read the entire input from the pipe
		inputBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}

		// Process the input using the processContent function
		processContent(inputBytes)
	}
}

func processFile(filename string) {
	fmt.Println("Processing file:", filename)
	// Add your file processing logic here

	// Read the content of the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		return
	}

	// Process the content using the processContent function
	processContent(content)
}

func processContent(content []byte) {
	// Convert content to lowercase and split into words
	text := strings.ToLower(string(content))
	wordRegex := regexp.MustCompile(`\b[\p{L}\p{Nd}'’]+\b`)
	words := wordRegex.FindAllString(text, -1)

	// Calculate three-word sequences and their frequencies
	trigramFreq := make(map[string]int)
	for i := 0; i <= len(words)-3; i++ {
		trigram := strings.Join(words[i:i+3], " ")
		trigramFreq[trigram]++
	}

	// Convert trigram frequencies to a slice for sorting
	trigramSlice := make([]trigramEntry, 0, len(trigramFreq))
	for trigram, freq := range trigramFreq {
		trigramSlice = append(trigramSlice, trigramEntry{Trigram: trigram, Frequency: freq})
	}

	// Sort trigram entries by frequency in descending order
	sortTrigramSlice(trigramSlice)

	// Print sorted trigram frequencies
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
