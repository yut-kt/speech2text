package htkdict

import (
	"bufio"
	"log"
	"strings"
)

func ScanWords(scanner *bufio.Scanner) []*Word {
	words := make([]*Word, 0)
	for scanner.Scan() {
		entry, output, phonemes := scanLine(scanner.Text())
		words = append(words, &Word{
			Entry:    entry,
			Output:   output,
			Phonemes: phonemes,
		})
	}
	return words
}

func scanLine(s string) (entry, output string, phonemes []string) {
	lineSlice := strings.Split(s, "\t")
	switch len(lineSlice) {
	case 3:
		entry = lineSlice[0]
		output = lineSlice[1]
		phonemes = strings.Split(lineSlice[2], " ")
	case 5:
		entry = lineSlice[0]
		output = lineSlice[3]
		phonemes = strings.Split(lineSlice[4], " ")
	default:
		log.Fatalf("ERROR: unknown slice size")
	}
	return
}
