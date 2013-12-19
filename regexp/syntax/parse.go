package main

import (
	"fmt"
	"regexp/syntax"
)

func main() {
	regex, _ := syntax.Parse(`[0120-2]@[ab][0-9]`, 'i')

	printSummary(regex)
}

func printSummary(regex *syntax.Regexp) {
	fmt.Printf("%v has %d sub expressions\n", regex, len(regex.Sub))

	for i, s := range regex.Sub {
		fmt.Printf("Child %d:\n", i)
		printSummary(s)
	}
}
