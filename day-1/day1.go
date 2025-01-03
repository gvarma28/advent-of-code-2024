package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	left := []int{}
	right := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")

		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	sort.Ints(left)
	sort.Ints(right)

	res := 0
	i := 0
	for {
		if i >= len(left) {
			break
		}

		diff := left[i] - right[i]
		if diff < 0 {
			diff *= -1
		}
		res += diff
		i++
	}

	fmt.Println(res)
}
