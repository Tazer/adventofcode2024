package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/tazer/adventofcode2024/utils"
)

func main() {
	list1, list2 := GetFileData(false)

	dist := GetDistances(list1, list2)

	fmt.Printf("Distance: %d\n", dist)

	match := GetMatchings(list1, list2)

	fmt.Printf("Match: %d\n", match)
}

func GetMatchings(list1, list2 []int) int {
	match := map[int]Pair{}

	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				if _, ok := match[i]; !ok {
					match[i] = Pair{Val: list1[i], Number: 1}
				} else {
					p := match[i]
					p.Number++
					match[i] = p
				}

			}
		}
	}

	finalResult := 0
	for _, v := range match {
		finalResult += v.Number * v.Val
	}
	return finalResult
}

type Pair struct {
	Val    int
	Number int
}

func GetDistances(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	dist := 0

	for i := 0; i < len(list1); i++ {
		addDist := list2[i] - list1[i]
		if addDist < 0 {
			addDist = list1[i] - list2[i]
		}
		fmt.Printf("list1[%d] = %d, list2[%d] = %d, addDist = %d\n", i, list1[i], i, list2[i], addDist)
		dist += addDist
	}

	return dist
}

func GetFileData(test bool) ([]int, []int) {
	var nums []int
	var nums2 []int

	utils.ReadFile(test, func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			row := scanner.Text()
			columns := strings.Split(row, "  ")
			num, err := strconv.Atoi(strings.TrimSpace(columns[0]))
			if err != nil {
				fmt.Println(err)
			}
			nums = append(nums, num)
			num2, err := strconv.Atoi(strings.TrimSpace(columns[1]))
			if err != nil {
				fmt.Println(err)
			}
			nums2 = append(nums2, num2)
		}

	})
	return nums, nums2
}
