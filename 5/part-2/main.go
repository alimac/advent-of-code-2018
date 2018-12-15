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
	str := ""
	for scanner.Scan() {
		str = scanner.Text()
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	az := []rune(alphabet)

	currentShort := len(str)
	pair := ""

	// Iterate over alphabet
	for _, r := range az {

		s := str

		// Pair
		char := string(r)
		upcase := strings.ToUpper(char)

		if strings.Contains(s, char) {

			// Remove pair from string
			s = strings.Replace(s, char, "", -1)
			s = strings.Replace(s, upcase, "", -1)

			// React polymer
			_, length := reactPolymer(s)

			// Track shortest polymer
			if length < currentShort {
				currentShort = length
				pair = char + upcase
			}
		}
	}
	fmt.Println(pair, currentShort)
}

func reactPolymer(s string) (string, int) {
	a := []rune(s)

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

	return string(a), len(a)
}
