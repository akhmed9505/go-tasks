package main

import "fmt"

/*
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.
Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента
(copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.
*/

func removeEl(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Некорректный индекс!")
		return slice
	}

	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = 0
	return slice[:len(slice)-1]
}

func removePtr(slice []*string, i int) []*string {
	if i < 0 || i >= len(slice) {
		fmt.Println("Некорректный индекс!")
		return slice
	}

	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = nil
	return slice[:len(slice)-1]
}

func main() {
	nums := []int{1, 2, 3}

	s1, s2, s3 := "a", "b", "c"
	ptrs := []*string{&s1, &s2, &s3}

	fmt.Println(removeEl(nums, 0))

	res := removePtr(ptrs, 0)
	for _, ptr := range res {
		fmt.Print(*ptr, " ")
	}
}
