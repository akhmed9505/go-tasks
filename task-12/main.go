package main

import "fmt"

/*
Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
Создать для неё собственное множество. Ожидается: получить набор уникальных слов.
Для примера, множество = {"cat", "dog", "tree"}.
*/

func makeSet(strs []string) []string {
	set := make(map[string]bool)
	for _, word := range strs {
		set[word] = true
	}

	res := make([]string, 0, len(set))
	for word := range set {
		res = append(res, word)
	}
	return res
}

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	res := makeSet(slice)
	fmt.Printf("Множество: %v\nКоличество уникальных слов: %d\n", res, len(res))
}
