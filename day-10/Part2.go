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

	part2(slic)
}

func part2(slic []string) {
	sprite := 1
	sum := 0
	j := 0
	k := 40
	for i := 1; i < 240; i += j {
		for j = i; sum < k; j++ {
			if slic[j] == "addx" && sum == sprite || slic[j] == "addx" && sum == sprite-1 || slic[j] == "addx" && sum == sprite+1 {
				fmt.Print("##")
				sum += 2

			} else if slic[j] == "addx" && sum == sprite || slic[j] == "addx" && sum == sprite-1 || slic[j] == "addx" && sum == sprite+1 {
				fmt.Print("##")
				sum += 2

			} else if slic[j] == "addx" && sum != sprite || slic[j] == "addx" && sum != sprite-1 || slic[j] == "addx" && sum != sprite+1 {
				fmt.Print("..")
				sum += 2
			} else if slic[j] == "noop" && sum == sprite || slic[j] == "noop" && sum == sprite-1 || slic[j] == "noop" && sum == sprite+1 {
				fmt.Print("#")
				sum++
			} else if slic[j] == "noop" && sum != sprite || slic[j] == "addx" && sum != sprite-1 || slic[j] == "addx" && sum != sprite+1 {
				fmt.Print(".")
				sum++
			} else {
				y, _ := strconv.Atoi(slic[j])
				sprite = sprite + y

			}
		}
		k += 40
		fmt.Println("")
	}
}
