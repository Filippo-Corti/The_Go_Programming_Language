package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)
 
var KEYWORDS = []string{"break", "case", "char", "const", "continue", "default", "defer", "else", "fallthrough", "for", "func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select", "struct", "switch", "type", "var"}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("a file name, please")
		os.Exit(0)
	}
	keywords, ok := findKeywordsInFile(os.Args[1])
	if !ok {
		fmt.Println("file not found")
		os.Exit(0)
	}
	printKeywords(keywords)
}

func findKeywordsInFile(fileName string) (map[string][]int, bool) {
	keywords := make(map[string][]int)
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return keywords, false
	}
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		updateKeywords(keywords, scanner.Text(), lineNumber)
		lineNumber++
	}
	return keywords, true
}

func updateKeywords(keywords map[string][]int, line string, lineNumber int)  {
	for _, keyword := range KEYWORDS {
		if strings.Contains(line, keyword) {
			keywords[keyword] = append(keywords[keyword], lineNumber)
		}
	}
}

func printKeywords(m map[string][]int) {
	keys := getKeys(m)
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %v\n", key, m[key])
	}
}

func getKeys(m map[string][]int) (keys []string) {
	for k, _ := range m {
		keys = append(keys, k)
	}
	return
}