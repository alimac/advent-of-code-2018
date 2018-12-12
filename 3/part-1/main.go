package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coordinates := make(map[string]int)

	for scanner.Scan() {
		claim := strings.Replace(scanner.Text(), ":", "", -1)
		parts := strings.Split(claim, " ")
		// parts[0] #1
		// parts[1] @

		start := strings.Split(parts[2], ",")
		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])

		dimensions := strings.Split(parts[3], "x")
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {

				point := fmt.Sprintf("%d,%d", startX+i, startY+j)
				coordinates[point]++
			}
		}
	}

	inches := 0
	for _, v := range coordinates {
		// More than one claim uses this coordinate
		if v > 1 {
			inches++
		}
	}
	fmt.Println(inches)
}
