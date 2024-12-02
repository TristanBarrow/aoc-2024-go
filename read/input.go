package read

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getContent(path string) []string {
    file, err := filepath.Abs(path)
    println(file)
	if err != nil {
		log.Panicf("Could not find file: %v", file)
		return nil
	}
	content, err := os.ReadFile(file)
	if err != nil {
		log.Panicf("Failed to read file: %v", file)
		return nil
	}
	return strings.Split(string(content), "\n")

}

func Day(num string) []string {
    return getContent("./input/day" + num + ".txt")
}

func Exm(num string) []string {
    return getContent("./input/example" + num + ".txt")
}
