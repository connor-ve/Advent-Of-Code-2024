package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	integers := CreateArray()
	// PartOne(integers)
	PartTwo(integers)
}

func CreateArray() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(0)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var nums [][]int
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		var temp []int
		for _, val := range numbers {
			intVal, _ := strconv.Atoi(val)
			temp = append(temp, intVal)
		}
		nums = append(nums, temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return nums
}

func PartOne(x [][]int) {
	ttl := 0
	for _, v := range x {
		asc := true
		good := true
		if v[0] > v[len(v)-1] {
			asc = false
		}
		var prev int
		fmt.Println(v, asc)
		for idx, val := range v {
			if idx == 0 {
				prev = val
				continue
			}
			if !checkDiff(prev, val, asc) {
				good = false
				break
			}
			prev = val
		}
		if good {
			ttl++
		}
	}
	fmt.Println("Part One Total :", ttl)
}

func checkAsc(x int, y int, z bool) bool {
	if z {
		if x < y {
			return true
		}
	} else {
		if x > y {
			return true
		}
	}
	return false
}

func checkDiff(x int, y int, z bool) bool {
	ascCheck := checkAsc(x, y, z)
	if !ascCheck {
		return false
	}
	diff := x - y
	if z {
		diff *= -1
	}
	if diff <= 3 && diff >= 1 {
		return true
	}
	return false
}

func PartTwo(x [][]int) {
	ttl := 0
	// v := x[0] // Debugging
	for _, v := range x {

		asc := true

		if v[0] > v[len(v)-1] {
			asc = false
		}

		fmt.Println(v, asc)

		for i := 0; i < len(v); i++ {
			temp := make([]int, 0, len(v)-1)
			temp = append(temp, v[:i]...)
			temp = append(temp, v[i+1:]...) 

			tempAsc := true
			if temp[0] > temp[len(temp)-1] {
				tempAsc = false
			}

			fmt.Println(temp)

			prev := temp[0]
			good := true
			for idx, val := range temp {
				if idx == 0 {
					continue
				}
				if !checkDiff(prev, val, tempAsc) {
					good = false
					break
				}
				prev = val
			}

			if good {
				ttl++
				break
			}
		}

	}
	fmt.Println("Part Two Total :", ttl)
}
