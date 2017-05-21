package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

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

	numArray := []int{10, 4, 5, 19, 30}
	i := searchAsc(numArray, 11)
	fmt.Println(numArray)
	fmt.Println(i)

	persons := []Person{
		Person{
			Name: "Hoge",
			Age:  10,
		},
		Person{
			Name: "Tom",
			Age:  5,
		},
		Person{
			Name: "Tom Jr",
			Age:  5,
		},
		Person{
			Name: "MaxBob",
			Age:  25,
		},
		Person{
			Name: "Jackson",
			Age:  7,
		},
	}
	sortSliceByName(persons)
	sortSliceByAge(persons)
}

func searchAsc(nums []int, x int) int {
	sort.Ints(nums)
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= x
	})
	return i
}

func sortSliceByName(persons []Person) {
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Name < persons[j].Name
	})
	fmt.Println("By Name: ", persons)
}

func sortSliceByAge(persons []Person) {
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age < persons[j].Age
	})
	fmt.Println("By Age: ", persons)
}
