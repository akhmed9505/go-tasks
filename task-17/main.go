package main

import "fmt"

/*
Реализовать алгоритм бинарного поиска встроенными методами языка.
Функция должна принимать отсортированный слайс и искомый элемент,
возвращать индекс элемента или -1, если элемент не найден.
Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
*/

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2 // безопасное вычисление середины (важно для малых типов)
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	fmt.Println(binarySearch(slice, target))
}
