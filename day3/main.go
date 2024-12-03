package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "os"
)

func main() {
	fmt.Println("Welcome to Day 3")
	PartOne(getChars())
	PartTwo(getChars())
}

func PartOne(s string) {
	r, _ := regexp.Compile(`mul\((\d*)?,(\d*)?\)`)

	matches := r.FindAllStringSubmatch(s, -1)

	ttl := 0
	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		Second, _ := strconv.Atoi(match[2])
		ttl += first * Second
	}
	fmt.Println(ttl)
}

func PartTwo(s string) {
	r, _ := regexp.Compile(`mul\((\d*)?,(\d*)?\)|do\(\)|don't\(\)`)

	matches := r.FindAllStringSubmatch(s, -1)

	ttl := 0
	enabled := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		}

		if enabled {
			first, _ := strconv.Atoi(match[1])
			Second, _ := strconv.Atoi(match[2])
			ttl += first * Second
		}
	}
	fmt.Println(ttl)
}

func getChars() string {
	filename := "input.txt"

	filebuffer, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	return inputdata
}
