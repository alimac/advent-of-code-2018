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

	frequency := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		change, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		frequency += change
	}
	fmt.Println(frequency)
}
