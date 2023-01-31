package main

import "fmt"
import "strings"

func main() {
	var url string
	fmt.Print("Inserisci URL: ")
	fmt.Scan(&url)
	fmt.Println("Host:", findHostInUrl(url))
	fmt.Println("Host:", findHostInUrlNoLibrary(url))
}

func findHostInUrl(s string) string {
	//Find index of "//" inside s
	indexOfDoubleSlash := strings.Index(s, "//") 
	//Find index of "/" in the substring that starts from first rune after "//"
	indexOfFirstSingleSlash := strings.Index(s[strings.Index(s, "//") + 2:], "/") 
	//if indexOfFirstSingleSlash is equal to -1, return the string from // to the end
	if indexOfFirstSingleSlash == -1 {
		return s[indexOfDoubleSlash + 2 :]
	}
	//Otherwise print the correct string between // and /
	return s[indexOfDoubleSlash + 2 : indexOfDoubleSlash + 2 + indexOfFirstSingleSlash]
}

func findHostInUrlNoLibrary(s string) string {
	var indiceDoppiSlash int
	for i := 0; i < len(s); i++ {
		if s[i] == '/' && s[i+1] == '/' {
			indiceDoppiSlash = i
			break;
		}
	}
	for i := indiceDoppiSlash + 2; i < len(s); i++ {
		if s[i] == '/' {
			return s[indiceDoppiSlash + 2 : i]
		}
	}
	return s[indiceDoppiSlash + 2 : ]
}