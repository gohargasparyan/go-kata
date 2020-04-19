package main

import "fmt"

func main() {
	s := []int{1, 2}
	median :=len(s)/2

	s = s[:median]
	fmt.Println(s)
}
