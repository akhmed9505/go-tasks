package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на вход строку.
Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).
*/

func reverse(str string) string {
	runes := []rune(str)
	i, j := 0, len(runes)-1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

func main() {
	fmt.Printf("главрыба\n%s", reverse("главрыба"))
}
