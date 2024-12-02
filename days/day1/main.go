package main

import (
	"bufio"
	"fmt"

	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Ziqiongg/advent-of-code-2024/utils"
)
func main() {
	var list1, list2 []int

    err := utils.ReadFile("input.txt", func(line string) {
        parts := strings.Fields(line)
        if len(parts) == 2 {
            num1, _ := strconv.Atoi(parts[0])
            num2, _ := strconv.Atoi(parts[1])
            list1 = append(list1, num1)
            list2 = append(list2, num2)
        }
    })

    if err != nil {
        fmt.Println("Error:", err)
    }
	result1 := partOne(list1, list2)
	fmt.Println(result1)
	result2 := partTwo(list1, list2)
	fmt.Println(result2)
}

func readFile() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				list1 = append(list1, num1)
				list2 = append(list2, num2)
			} else {
				fmt.Println("Error converting string to int:", err1, err2)
			}
		}
	}
	return list1, list2
}

func partOne(list1, list2 []int) int{
	sort.Ints(list1)
	sort.Ints(list2)
	var result int
	for i, num := range list1 {
		result += utils.Abs(num, list2[i])
	}

	return result
}

func partTwo(list1, list2 []int) int {
	var result int

	// convert list2 to map
	mapTwo := make(map[int]int)
	for _, num := range list2 {
		mapTwo[num] += 1
	}

	//convert list1 to set
	setOne := make(map[int]struct{})
	for _, num := range list1 {
		setOne[num] = struct{}{}
	}

	for num, _ := range setOne {
		times := mapTwo[num]
		result += times * num
	}

	return result
}