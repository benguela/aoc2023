package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	max_cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	tot := 0
	for _, line := range fileLines {
		games := strings.Split(line, ": ")
		id, err := strconv.Atoi(strings.Split(games[0], " ")[1])
		if err != nil {
			// ... handle error
			panic(err)
		}

		sets := strings.Split(games[1], "; ")

		valid := true

	out:
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				cubit := strings.Split(cube, " ")
				i, err := strconv.Atoi(cubit[0])
				if err != nil {
					panic(err)
				}
				if i > max_cubes[cubit[1]] {
					valid = false
					break out
				}
			}
		}

		if valid {
			tot += id
		}
	}
	fmt.Println(tot)
}
