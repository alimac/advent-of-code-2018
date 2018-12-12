package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id     string
	startX int
	startY int
	height int
	width  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coordinates := [1000][1000]int{}
	claims := []claim{}

	for scanner.Scan() {
		str := strings.Replace(scanner.Text(), ":", "", -1)
		str = strings.Replace(str, "#", "", -1)
		parts := strings.Split(str, " ")

		start := strings.Split(parts[2], ",")
		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])

		dimensions := strings.Split(parts[3], "x")
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])

		//
		claims = append(claims, claim{
			id:     parts[0],
			startX: startX,
			startY: startY,
			height: height,
			width:  width,
		})

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				coordinates[startX+i][startY+j]++
			}
		}
	}

	// For a claim that doesn't overlap other claims
	// adding up coordinates should equal length * width
	for _, c := range claims {
		sum := 0
		for i := c.startX; i < c.width+c.startX; i++ {
			for j := c.startY; j < c.height+c.startY; j++ {
				sum += coordinates[i][j]
			}
		}
		if sum == c.height*c.width {
			fmt.Println(c.id)
			return
		}
	}
}
