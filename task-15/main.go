package main

import (
	"fmt"
	"strings"
)

/*
Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?
Приведите корректный пример реализации.

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

Вопрос: что происходит с переменной justString?
*/

func createHugeString(n int) string {
	return strings.Repeat("a", n)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
	fmt.Println(justString)
}

/*
В старом способе подстрока ссылалась на весь буфер v (1024 байта), хотя нам нужно было только первые 100 байт.
Это приводило к утечке памяти, так как GC не мог очистить буфер, пока жива переменная justString.

В новом коде мы создаём строку string([]byte(v[:100])), сначала копируем только нужные 100 байт в
новый массив и затем создаем из него строку.

Теперь justString содержит только нужную часть данных, а оригинальный буфер v с
лишними данными может быть очищен сборщиком мусора.
*/
