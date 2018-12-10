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
	var ids []string

	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	// Iterate over ID pairs
	for _, id1 := range ids {
		for _, id2 := range ids {

			// Skip if its the same ID
			if id1 == id2 {
				continue
			}

			// Track number of differences
			diff := 0
			// Track position where difference occurs
			position := 0

			for i := range id1 {
				// Increment if there is a difference
				if id1[i] != id2[i] {
					diff++
					position = i
				}
				// Skip there is more than 1 difference
				if diff > 1 {
					continue
				}
			}
			// We found it!!
			if diff == 1 {
				// Concatenate everything before the "difference position" and everything after
				id := id1[:position] + id1[position+1:]
				fmt.Printf("%s", id)
				return
			}
		}
	}
}
