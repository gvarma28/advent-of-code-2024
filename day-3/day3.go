package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	scanner.Scan()
	input := scanner.Text()

	pattern := `mul\(\d+,\d+\)`
	reg, _ := regexp.Compile(pattern)

	matches := reg.FindAllString(input, -1)

	res := 0
	for _, value := range matches {
		first := strings.Split(strings.Split(value, "(")[1], ",")[0]
		second := strings.Split(strings.Split(value, ",")[1], ")")[0]

		if len(first) > 3 || len(second) > 3 {
			continue
		}

		firstNum, _ := strconv.Atoi(first)
		secondNum, _ := strconv.Atoi(second)

		res += firstNum * secondNum
	}

	fmt.Println(res)

}

// func solution_without_regex() {
// 	// open the input file
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println("Error while opening the file: ", err)
// 		return
// 	}
// 	defer file.Close()
//
// 	// scanner to read the file line by line
// 	scanner := bufio.NewScanner(file)
//
// 	scanner.Scan()
// 	input := scanner.Text()
//
// 	n := len([]rune(input))
// 	p := 0
// 	for {
// 		if p >= n-4 {
// 			break
// 		}
// 		if input[p:p+4] == "mul(" {
// 			q := p + 4
//
// 			for {
// 				if input[q] == ")" {
//
// 				}
//
// 			}
//
// 			fmt.Println(input[p : p+10])
// 		}
// 		p++
// 	}
// }
//
