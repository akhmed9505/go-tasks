package main

import "fmt"

/*
Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
Для выбора опорного элемента можно взять середину или первый элемент.
*/

func quickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	pivot := slice[0]
	var low, high []int

	for _, num := range slice[1:] {
		if num < pivot {
			low = append(low, num)
		} else {
			high = append(high, num)
		}
	}
	return append(append(quickSort(low), pivot), quickSort(high)...)
}

func main() {
	nums := []int{5, 3, 6, 8, 4, 9, 2, 7, 1, 10}
	fmt.Println("Исходный:", nums)
	fmt.Println("Отсортированный:", quickSort(nums))
}
