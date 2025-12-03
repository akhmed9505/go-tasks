package main

import (
	"fmt"
	"math"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить эти значения в группы с шагом 10 градусов.
Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.
Пояснение: диапазон -20 включает значения от -20 до -29.9, диапазон 10 – от 10 до 19.9, и т.д.
Порядок в подмножествах не важен.
*/

func getGroup(temp float64) float64 {
	if temp < 0 {
		return math.Ceil(temp/10) * 10
	} else {
		return math.Floor(temp/10) * 10
	}
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[float64][]float64)

	for _, temp := range temps {
		group := getGroup(temp)
		groups[group] = append(groups[group], temp)
	}

	fmt.Println(groups)
}
