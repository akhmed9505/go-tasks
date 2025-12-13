package main

import "fmt"

/*
Разработать программу, которая переворачивает порядок слов в строке.
Пример: входная строка:
«snow dog sun», выход: «sun dog snow».
Считайте, что слова разделяются одиночным пробелом.
Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
*/

// Функция для разворота всей строки
func reverse(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

// Функция для разворота слов
func reverseWords(str string) string {
	runes := []rune(str)
	reverse(runes, 0, len(runes)-1) // Разворачиваем всю строку

	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	str := "snow dog sun"
	fmt.Printf("Before: %s\nAfter: %s", str, reverseWords(str))
}
