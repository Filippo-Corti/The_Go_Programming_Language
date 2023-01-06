package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

//It replaces the n-th instance of old with new, inside s
//It does so by getting the index of the instance to replace and calling the strings.Replace function
//on the substring starting from index
//If the function getNthIndexOfOld returns -1, it returns the starting string
func sostituisci(s, old, new []byte, n int) []byte {
	indexOfOld := getNthIndexOfOld(s, old, n)
	if indexOfOld == -1 {
		return s
	}
	return []byte(string(s[:indexOfOld]) + strings.Replace(string(s[indexOfOld:]), string(old), string(new), 1))
}

//Returns the index of the n-th instance of old inside s
//It does so by repeating the function strings.Index n times on substrings of s
//If the n-th instance doesn't exist, it returns -1
func getNthIndexOfOld(s, old []byte, n int) (index int) {
	sString := string(s)
	oldString := string(old)
	index = strings.Index(sString, oldString)
	for i := 1; i < n; i++ {
		relIndex := strings.Index(sString[index + len(old):], oldString)
		if relIndex == -1 {
			return -1
		}
		index += relIndex + len(old)
	}
	return
}

func main() {
	n, _ := strconv.Atoi(os.Args[4])
	fmt.Println(os.Args[1])
	fmt.Println(string(sostituisci([]byte(os.Args[1]), []byte(os.Args[2]), []byte(os.Args[3]), n)))
}