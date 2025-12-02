package main

import "fmt"

/*
Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
Подсказка: используйте битовые операции (|, &^).
*/

func SetBit(num int64, position, value int) int64 {
	if value == 0 {
		return num &^ (int64(1) << position)
	}
	return num | (int64(1) << position)
}

func main() {
	var num int64           // число, которое будет изменено
	var position, value int // позиция i-го бита, значение 1 или 0

	fmt.Println("Введите число, позицию i-го бита, и значение 1 или 0")
	fmt.Scan(&num, &position, &value)
	if value != 0 && value != 1 {
		fmt.Println("Ошибка: значение должно быть 0 или 1")
		return
	}

	result := SetBit(num, position, value)
	fmt.Printf("Результат: %d (двоичное: %b)\n", result, result)
}
