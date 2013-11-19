/*
MustCompile() function lets you create a regexp and assign it to a var at the package level
SubMatch() methods return a []string indexed by the match group position: 0 -> entire match 
*/

package main

import (
	"fmt"
	"regexp"
)

// var digitsRegexp = regexp.MustCompile(`\d+`)
var digitsRegexp = regexp.MustCompile(`(\d+)\D+(\d+)`)

func main() {
	someString := "#1000ab*cd123and456x"

        // entire match, 1st, 2nd: 1000ab*cd123 1000 123
        fmt.Println(digitsRegexp.FindStringSubmatch(someString))

	// Find just the leftmost
	fmt.Println(digitsRegexp.FindString(someString))

	// Find all (-1) the matches
	fmt.Println(digitsRegexp.FindAllString(someString, -1))
}

