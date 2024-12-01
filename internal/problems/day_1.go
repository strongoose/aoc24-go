package problems

import (
	"sort"
	"strconv"
	"strings"
)

func parse_day1_input(input string) ([]int, []int) {
	input = strings.TrimSpace(input)

	var left, right []int

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "   ")

		ln, _ := strconv.Atoi(parts[0])
		rn, _ := strconv.Atoi(parts[1])

		left = append(left, ln)
		right = append(right, rn)
	}

	return left, right
}

func Day1Pt1(input string) int {

	left, right := parse_day1_input(input)

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i, ln := range left {
		sum += Abs(ln - right[i])
	}

	return sum
}

func Day1Pt2(input string) int {
	left, right := parse_day1_input(input)

	similarity := func(n int, xs []int) int {
		count := 0
		for _, x := range xs {
			if x == n {
				count += 1
			}
		}

		return n * count
	}

	sum := 0
	for _, ln := range left {
		sum += similarity(ln, right)
	}

	return sum
}
