package main

import (
	"fmt"
	"os"

	"github.com/strongoose/aoc24-go/internal/problems"
)

func read_input(day uint8) string {
	input_file := fmt.Sprintf("input/2024/%d.txt", day)
	contents, err := os.ReadFile(input_file)
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [day number]\n", os.Args[0])
	os.Exit(1)
}

func main() {

	if len(os.Args) != 2 {
		usage()
	}

	problem := os.Args[1]

	switch problem {
	case "1":
		input := read_input(1)
		fmt.Printf("Part 1: %d\n", problems.Day1Pt1(input))
		fmt.Printf("Part 2: %d\n", problems.Day1Pt2(input))
	}
}
