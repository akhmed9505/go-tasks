package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/

func main() {
	var wg sync.WaitGroup
	array := [5]int{2, 4, 6, 8, 10}
	result := make([]int, len(array))

	for i, v := range array {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// В Go 1.22 проблема гонки данных (Data Race), связанная с переменными цикла, была автоматически устранена.
			result[i] = v * v
		}()
	}

	wg.Wait()
	for _, v := range result {
		fmt.Println(v)
	}
}
