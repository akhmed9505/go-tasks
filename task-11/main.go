package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) —
т.е. вывести элементы, присутствующие и в первом, и во втором.
Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3}
*/

func makeSet(slice []int) map[int]bool {
	set := make(map[int]bool)
	for _, num := range slice {
		set[num] = true
	}
	return set
}

func intersection(setA, setB map[int]bool) []int {
	var res []int
	for num := range setA {
		if _, ok := setB[num]; ok {
			res = append(res, num)
		}
	}
	return res
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	setA := makeSet(A)
	setB := makeSet(B)

	fmt.Println("Пересечение:", intersection(setA, setB))
}
