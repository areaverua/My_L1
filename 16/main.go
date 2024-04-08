package main

import (
	"fmt"
	"sort"
)
// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.




func main() {

	nums := []int{1, 8, 2, 1, 4, 7, 3, 2, 3, 6}
	quickSort(nums, 0, len(nums)) //   первая реализация
	fmt.Println(nums)

	a := []int{-1, 2, 1, 4, 23, 0, -12, 129}
	sort.Ints(a) // вторая реализация
	fmt.Println(a)

}
func quickSort(a []int, l int, r int) {
	i := l
	j := r - 1
	mid := a[(i+j)/2]

	for i <= j {
		for a[i] < mid {
			i++
		}
		for a[j] > mid {
			j--
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	if j > l {
		quickSort(a, l, j+1)

	}

	if i < r {
		quickSort(a, i, r)
	}
}
