package main

import (
	"aoc/days"
	"log"
	"os"
	"strings"
)

const root = "/Users/tristanbarrow/Projects/aoc_go"

func main() {
	// println(days.Day1Part1(readFile("day1")))
    println(days.Day1Part2(readFile("day1")))
}

func readFile(file string) []string {
	filename := root + "/input/" + file + ".txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Panicf("Failed to read file: %v", filename)
		return nil
	}
	return strings.Split(string(content), "\n")
}
