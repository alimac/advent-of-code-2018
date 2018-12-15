package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var a []rune
	for scanner.Scan() {
		s := scanner.Text()
		a = []rune(s)
	}

	for i := 0; i < len(a); i++ {
		// Figure out next position
		next := i + 1
		// If we are not at array length and runes differ
		if next < len(a) && a[i] != a[next] {
			// Convert to lowercase - if they match, we found polarity!
			if strings.ToLower(string(a[i])) == strings.ToLower(string(a[next])) {
				// Remove both runes
				a = append(a[:i], a[i+2:]...)
				// Decrement index by 2
				if i > 1 {
					i -= 2
					// Unless we are at the beginning
					// Set to -1 to start a zero (++)
				} else {
					i = -1
				}
			}
		}
	}

	fmt.Println(string(a), len(a))
}
