package main

import (
	"fmt"

	mq "github.com/CodingDog-CD/sorting/pkg/merge_sort"
	qc "github.com/CodingDog-CD/sorting/pkg/quick_sort"
)

func main() {
	input := []int{65, 18, 34, 156, 48, 64, 30, 27}
	fmt.Println("----- Quick Sort -----")
	fmt.Printf("Before sort: %v\n", input)
	qc.Sort(input, 0, len(input)-1)
	fmt.Printf("After sort: %v\n", input)
	fmt.Println("----- voila! -----")

	input = []int{65, 18, 34, 156, 48, 64, 30, 27}
	fmt.Println("----- Merge Sort -----")
	fmt.Printf("Before sort: %v\n", input)
	sortedArray := mq.Sort(input)
	fmt.Printf("After sort: %v\n", sortedArray)
	fmt.Println("----- voila! -----")

}
