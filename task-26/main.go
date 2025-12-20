package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке встречаются
один раз (т.е. строка состоит из уникальных символов).
Вывод: true, если все символы уникальны, false, если есть повторения.
Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.
Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.
Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/

func isUniqueChars(s string) bool {
	mp := make(map[rune]bool)
	s = strings.ToLower(s)

	for _, char := range s {
		if mp[char] {
			return false
		}
		mp[char] = true
	}
	return true
}

func main() {
	fmt.Println(isUniqueChars("abcd"))
	fmt.Println(isUniqueChars("abCdefAaf"))
	fmt.Println(isUniqueChars("aabcd"))
}
