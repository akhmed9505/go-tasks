package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.
Подсказка: используйте time.After или таймер для ограничения времени работы.
*/

func reader(ch chan int, done chan struct{}) {
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				return
			}
			time.Sleep(500 * time.Millisecond) // Симуляция обработки сообщения
			fmt.Println("Получено:", value)
		case <-done:
			return
		}
	}
}

func main() {
	var n int
	fmt.Printf("Введите время выполнения программы в секундах:\n")
	fmt.Scan(&n)

	ch := make(chan int)
	done := make(chan struct{})

	go reader(ch, done)

	timer := time.After(time.Duration(n) * time.Second)

	i := 1
	for {
		select {
		case <-timer:
			close(done)
			fmt.Printf("Программа завершилась, прошло %d секунд", n)
			return
		case ch <- i:
			i++
		}
	}
}
