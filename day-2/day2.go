package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening the file: ", err)
		return
	}
	defer file.Close()

	// scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		strLevels := strings.Split(line, " ")

		level := []int{}
		for i := 0; i < len(strLevels); i++ {
			numLevel, _ := strconv.Atoi(strLevels[i])
			level = append(level, numLevel)
		}

		if len(level) <= 1 {
			res++
			continue
		}

		isIncreasing := true
		valid := true
		for i := 0; i < len(level)-1; i++ {
			if i == 0 {
				if level[i] > level[i+1] {
					isIncreasing = false
				}
				if !isValidRange(level[i], level[i+1]) {
					valid = false
					break
				}
				continue
			}

			if !isValidRange(level[i], level[i+1]) {
				valid = false
				break
			}

			if isIncreasing && level[i] > level[i+1] {
				valid = false
				break
			}

			if !isIncreasing && level[i] < level[i+1] {
				valid = false
				break
			}
		}

		if valid {
			res++
		}

	}
	fmt.Println(res)
}

func isValidRange(a int, b int) bool {
	diff := a - b
	if diff < 0 {
		diff *= -1
	}
	if diff >= 1 && diff <= 3 {
		return true
	}
	return false
}
