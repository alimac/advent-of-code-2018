package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	logs := []string{}

	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}
	sort.Strings(logs)
	currentID := 0
	falls := 0
	var minutes = map[int]map[int]int{}

	for _, l := range logs {
		// Set guard ID
		if strings.Contains(l, "Guard") {
			shift := strings.SplitAfter(l, "#")[1]
			idString := strings.Split(shift, " ")[0]
			currentID, err = strconv.Atoi(idString)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}

		if currentID > 0 {
			if strings.Contains(l, "falls") {
				falls, _ = strconv.Atoi(l[15:17])
			}

			if strings.Contains(l, "wakes") {
				wakes, _ := strconv.Atoi(l[15:17])

				for i := falls; i < wakes; i++ {
					// Gotta initialize map
					if minutes[currentID] == nil {
						minutes[currentID] = map[int]int{}
					}
					minutes[currentID][i]++
				}
			}
		}
	}

	guard := 0
	highest := 0
	min := 0

	for id, m := range minutes {
		for k, v := range m {
			if v > highest {
				highest = v
				min = k
				guard = id
			}
		}
	}
	fmt.Println(guard * min)
}
