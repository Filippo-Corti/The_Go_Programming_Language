package main

import "fmt"
import "bufio"
import "os"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Inserito %s\n", line)
	}
}