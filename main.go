package main

import (
	"fmt"
	"sort"
	"strconv"
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

	// search desc
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] <= 4
	})
	fmt.Println("index:" + strconv.Itoa(index))
	fmt.Println("value:" + strconv.Itoa(nums[index]))

	// search asc
	sort.Ints(nums)
	index = sort.Search(len(nums), func(i int) bool {
		return nums[i] >= 8
	})
	fmt.Println(nums)
	fmt.Println("index:" + strconv.Itoa(index))
	fmt.Println("value:" + strconv.Itoa(nums[index]))
}
