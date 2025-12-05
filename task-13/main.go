package main

import "fmt"

/*
Поменять местами два числа без использования временной переменной.
Подсказка: примените сложение/вычитание или XOR-обмен.
*/

func swapNums(x, y int) (int, int) {
	fmt.Println("До:", x, y)
	x = x ^ y
	fmt.Println("Шаг 1:", x, y)
	y = x ^ y
	fmt.Println("Шаг 2:", x, y)
	x = x ^ y
	fmt.Println("Шаг 3:", x, y)

	return x, y
}

func swapNumsAddSub(x, y int) (int, int) {
	fmt.Println("До:", x, y)
	x = x + y
	fmt.Println("Шаг 1:", x, y)
	y = x - y
	fmt.Println("Шаг 2:", x, y)
	x = x - y
	fmt.Println("Шаг 3:", x, y)

	return x, y
}

func main() {
	var x, y int
	fmt.Println("Введите 2 числа, которые хотите поменять местами: ")
	fmt.Scan(&x, &y)

	fmt.Println("\n--- XOR обмен ---")
	resX, resY := swapNums(x, y)
	fmt.Println("\nРезультат:", resX, resY)

	fmt.Println("\n--- Сложение/вычитание обмен ---")
	resX, resY = swapNumsAddSub(x, y)
	fmt.Println("\nРезультат:", resX, resY)
}
