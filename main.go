package main

import (
	"fmt"

	qc "github.com/CodingDog-CD/sorting/pkg/quick_sort"
)

func main() {
	input := []int{65, 18, 34, 156, 48, 64, 30, 27}
	fmt.Printf("Before sort: %v\n", input)
	qc.Sort(input, 0, len(input)-1)
	fmt.Printf("After sort: %v\n", input)
	fmt.Println("----- voila! -----")
}
