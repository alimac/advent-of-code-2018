package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nums []int

	seen := map[int]bool{0: true}
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		// var n int
		// _, err := fmt.Sscanf(scanner.Text(), "%d", &n)
		// if err != nil {
		// 	log.Fatalf("could not read %s: %v", scanner.Text(), scanner.Err())
		// }

		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}

	// Loop over array continuously
	for {
		for _, n := range nums {
			sum += n
			if seen[sum] {
				fmt.Println(sum)
				return
			}
			seen[sum] = true
		}
	}
}
