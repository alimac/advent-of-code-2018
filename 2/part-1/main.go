package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	twiceTotal := 0
	thriceTotal := 0

	for scanner.Scan() {
		seen := make(map[rune]int)
		twice := false
		thrice := false

		for _, char := range scanner.Text() {
			seen[char]++
		}

		for _, val := range seen {
			if val == 3 && !thrice {
				thriceTotal++
				thrice = true
			}
			if val == 2 && !twice {
				twiceTotal++
				twice = true
			}
		}
	}
	fmt.Println(twiceTotal * thriceTotal)
}
