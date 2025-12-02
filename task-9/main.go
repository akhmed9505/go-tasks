package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива,
во второй – результат операции x*2. После этого данные из второго канала должны выводиться в stdout.
То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
Убедитесь, что чтение из второго канала корректно завершается.
*/

func generate(ch1 chan int, nums []int) {
	for _, el := range nums {
		ch1 <- el
	}
	close(ch1)
}

func handle(ch1, ch2 chan int) {
	for value := range ch1 {
		ch2 <- 2 * value
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	nums := []int{10, 20, 30, 40, 50}

	go generate(ch1, nums)
	go handle(ch1, ch2)

	for value := range ch2 {
		fmt.Println(value)
	}
}
