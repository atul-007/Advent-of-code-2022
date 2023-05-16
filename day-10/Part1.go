package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	slic := make([]string, 1)
	for scanner.Scan() {
		slic = append(slic, scanner.Text())

	}
	part1(slic)

}

func part1(slic []string) {
	cycles := 0
	sum := 1
	sum2 := 0
	x := 20
	for i := 1; i < len(slic); i++ {
		if slic[i] == "noop" {
			cycles++
		}
		if cycles == x {
			sum2 = sum2 + (cycles * sum)
			x += 40
		}

		if slic[i] == "addx" {

			cycles++

		}
		if cycles == x {
			sum2 = sum2 + (cycles * sum)
			x += 40
		}
		if slic[i] == "addx" {

			cycles++

		}
		if cycles == x {
			sum2 = sum2 + (cycles * sum)
			x += 40
		}
		if slic[i] != "noop" && slic[i] != "addx" {
			y, _ := strconv.Atoi(slic[i])
			sum = sum + y

		}
		if cycles > 220 {
			break
		}
	}

	fmt.Println(sum2)

}
