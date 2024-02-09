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
		// re := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
		// nums := re.FindAllString(line, -1)
		//sneaky problem setters made overlapping string so now doing a brute force
		//not the most elegant solution and lots of converting between str and ints.
		var numbers []string
		for i := 0; i < len(line); i++ {
			n, err := strconv.Atoi(string(line[i]))
			if err != nil {
				if strings.Index(string(line[i:]), "one") == 0 {
					n = 1
				} else if strings.Index(string(line[i:]), "two") == 0 {
					n = 2
				} else if strings.Index(string(line[i:]), "three") == 0 {
					n = 3
				} else if strings.Index(string(line[i:]), "four") == 0 {
					n = 4
				} else if strings.Index(string(line[i:]), "five") == 0 {
					n = 5
				} else if strings.Index(string(line[i:]), "six") == 0 {
					n = 6
				} else if strings.Index(string(line[i:]), "seven") == 0 {
					n = 7
				} else if strings.Index(string(line[i:]), "eight") == 0 {
					n = 8
				} else if strings.Index(string(line[i:]), "nine") == 0 {
					n = 9
				}
			}
			if n != 0 {
				numbers = append(numbers, strconv.Itoa(n))
			}
		}
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
