package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых
переменных a, b, значения которых > 2^20 (больше 1 миллион).
Комментарий: в Go тип int справится с такими числами, но обратите внимание
на возможное переполнение для ещё больших значений. Для очень больших чисел можно использовать math/big.
*/

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1048577", 10)
	b.SetString("2097153", 10)

	addition := new(big.Int).Add(a, b)
	subtraction := new(big.Int).Sub(a, b)
	multi := new(big.Int).Mul(a, b)
	quot := new(big.Int)
	rem := new(big.Int)
	quot.DivMod(a, b, rem)

	fmt.Printf("a = %s\n", a)
	fmt.Printf("b = %s\n", b)
	fmt.Printf("Сумма: %s\n", addition)
	fmt.Printf("Разность: %s\n", subtraction)
	fmt.Printf("Произведение: %s\n", multi)
	fmt.Printf("Частное: %s\nОстаток: %s\n", quot, rem)
}
