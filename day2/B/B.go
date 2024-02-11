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

	tot := 0
	for _, line := range fileLines {
		games := strings.Split(line, ": ")

		sets := strings.Split(games[1], "; ")

		red := 0
		green := 0
		blue := 0
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				cubit := strings.Split(cube, " ")
				i, err := strconv.Atoi(cubit[0])
				if err != nil {
					panic(err)
				}
				switch cubit[1] {
				case "red":
					if i > red {
						red = i
					}
				case "green":
					if i > green {
						green = i
					}
				case "blue":
					if i > blue {
						blue = i
					}
				}
			}
		}

		tot += red * green * blue
	}
	fmt.Println(tot)
}
