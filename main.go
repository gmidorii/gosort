package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 2, 10, 8}
	sort.Ints(nums)
	fmt.Println(nums)
	// check sorted
	fmt.Println(sort.IntsAreSorted(nums))

	// sort reverse
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
}
