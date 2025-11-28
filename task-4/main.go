package main

import (
	"context"
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
Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.
*/

func worker(id int, ch <-chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершается, потому что получил сигнал отмены(cancel()).\n", id)
			return
		case msg, ok := <-ch:
			if !ok {
				fmt.Printf("Воркер %d завершается, потому что канал закрыт.\n", id)
				return
			}

			fmt.Printf("Воркер %d получил сообщение: %d\n", id, msg)
			// Симуляция обработки сообщения
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Пожалуйста, введите число воркеров. Пример использования: go run main.go 2")
	}

	nWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || nWorkers <= 0 {
		log.Fatal("Пожалуйста введите правильное значение для количества воркеров")
	}

	var wg sync.WaitGroup
	ch := make(chan int)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg, ctx)
	}

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	msgID := 1
	running := true

	for running {
		select {
		case <-ctx.Done():
			fmt.Println("Получен сигнал остановки(Context.Done()). Остановка...")
			running = false
		case <-ticker.C:
			select {
			case ch <- msgID:
				msgID++
			case <-ctx.Done():
				running = false // Гарантирует мгновенную реакцию на Ctrl+C
			}
		}
	}

	close(ch)
	fmt.Println("Канал сообщений закрыт. Ждём завершения воркеров...")

	wg.Wait()
	fmt.Println("Все воркеры закончили.")
}
