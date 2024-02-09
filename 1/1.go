package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	readFile, err := os.Open("input.txt")

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
		re := regexp.MustCompile("[0-9]")
		numbers := re.FindAllString(line, -1)
		cal := numbers[0] + numbers[len(numbers)-1]
		i, err := strconv.Atoi(cal)
		if err != nil {
			// ... handle error
			panic(err)
		}
		tot += i
	}

	fmt.Println(tot)

}
