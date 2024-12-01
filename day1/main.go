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
	PartOne()
	PartTwo()
}

// ! Successful answer
func PartOne() {
	NumLeft, NumRight := HandleFile()
	ttl := 0
	for i, _ := range NumLeft {
		val := NumLeft[i] - NumRight[i]
		if val < 0 {
			val *= -1
		}
		ttl += val
	}
	fmt.Println("Total Distance : ", ttl)
}

// ! Successful answer
func PartTwo() {
	NumLeft, NumRight := HandleFile()
	simScore := 0
	for _, val := range NumLeft {
		inc := 0
		for _, vl := range NumRight {
			if vl > val {
				break
			}
			if vl == val {
				inc++
			}
		}

		simScore += (val * inc)
	}
	fmt.Println("Similarity Score :", simScore)
}

// Pull data and sort int arrays 
func HandleFile() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(0)
	}
	defer file.Close()
	var NumLeft []int
	var NumRight []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		NumVal := strings.Fields(line)
		leftInt, _ := strconv.Atoi(NumVal[0])
		RightInt, _ := strconv.Atoi(NumVal[1])
		NumLeft = append(NumLeft, leftInt)
		NumRight = append(NumRight, RightInt)
	}
	sort.Ints(NumLeft)
	sort.Ints(NumRight)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return NumLeft, NumRight
}
