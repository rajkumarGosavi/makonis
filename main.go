package main

import (
	"fmt"
	"makonis/permutation"
)

func main() {
	s1 := "teststring"
	s2 := "stringtets"
	if permutation.CheckPermutationAscii(s1, s2) {
		printer("%s is a permutation of %s\n", s2, s1)
	} else {
		printer("%s is not a permutation of %s\n", s2, s1)
	}
}

func printer(format, s1, s2 string) {
	fmt.Printf(format, s1, s2)
}
