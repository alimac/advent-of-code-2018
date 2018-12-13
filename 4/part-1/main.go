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

	totalSleep := map[int]int{}
	var minutes = map[int]map[int]int{}
	currentTotal := 0
	sleepiestGuardID := 0

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
				sleep := wakes - falls

				totalSleep[currentID] += sleep

				if totalSleep[currentID] > currentTotal {
					sleepiestGuardID = currentID
					currentTotal = totalSleep[currentID]
				}

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

	min := 0
	highest := 0

	for k, v := range minutes[sleepiestGuardID] {
		if v > highest {
			highest = v
			min = k
		}
	}
	fmt.Println(sleepiestGuardID * min)
}
