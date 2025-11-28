package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (в главной горутине).
Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
*/

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("Воркер %d получил сообщение: %d\n", id, msg)
		// Симуляция обработки сообщений
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Воркер %d закончил.\n", id)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Пожалуйста, введите число воркеров. Пример использования: go run main.go 2")
	}

	var wg sync.WaitGroup

	ch := make(chan int)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	nWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || nWorkers <= 0 {
		log.Fatal("Пожалуйста введите правильное значение для количества воркеров")
	}

	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	msgID := 1
	running := true

	for running {
		select {
		case <-stop:
			fmt.Println("Получен сигнал остановки. Остановка...")
			running = false
		case <-ticker.C:
			ch <- msgID
			msgID++
		}
	}
	close(ch)
	fmt.Println("Канал сообщений закрыт. Ждём завершения воркеров...")

	wg.Wait()
	fmt.Println("Все воркеры закончили.")
}
