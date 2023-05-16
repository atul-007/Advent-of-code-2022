package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point2 struct {
	x, y int
}

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	visitedByTail := make(map[point2]bool)
	head := point2{0, 0}
	tail := point2{0, 0}
	visitedByTail[tail] = true
	for sc.Scan() {
		direction := rune(sc.Text()[0])
		moves, _ := strconv.Atoi(sc.Text()[2:])

		//I calculate moves one by one
		for moves > 0 {
			switch direction {
			case 'U':
				head.y++
			case 'R':
				head.x++
			case 'D':
				head.y--
			case 'L':
				head.x--
			}
			moves--
			tail = adjustTail(tail, head)
			visitedByTail[tail] = true
		}
	}

	fmt.Println(len(visitedByTail))
}

func adjustTail(tail point2, head point2) (newTail point2) {
	newTail = tail
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{-2, 1}, point2{-1, 2}, point2{0, 2}, point2{1, 2}, point2{2, 1}:
		newTail.y++
	}
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{1, 2}, point2{2, 1}, point2{2, 0}, point2{2, -1}, point2{1, -2}:
		newTail.x++
	}
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{2, -1}, point2{1, -2}, point2{0, -2}, point2{-1, -2}, point2{-2, -1}:
		newTail.y--
	}
	switch (point2{head.x - tail.x, head.y - tail.y}) {
	case point2{-1, -2}, point2{-2, -1}, point2{-2, -0}, point2{-2, 1}, point2{-1, 2}:
		newTail.x--
	}
	return
}
