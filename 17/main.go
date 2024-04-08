package main

import (
	"fmt"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка.
func main() {
	nums := []int{8, 2, 1, 4, 7, 3, 6}
	sort.Ints(nums)
	fmt.Println(nums)

	f := bin(nums, 7)
	fmt.Println(nums[f])

	a := []int{1, 2, 4, 7, 0, -123, 19}
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(sort.SearchInts(a, 2)) // Реализация с помощью пакета sort

}

func bin(slice []int, target int) int {
	left := 0
	right := len(slice) - 1

	for left <= right {
		mid := (left + right) / 2

		if slice[mid] < target {
			left = mid + 1
		} else if slice[mid] > target {
			right = mid - 1
		} else {
			return mid // Возращаем индекс target
		}

	}

	return 0

}
