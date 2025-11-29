package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.
Продемонстрируйте каждый способ в отдельном фрагменте кода.
*/

func main() {
	// ================================================
	//                Выход по условию
	// ================================================
	stop := false
	go func() {
		for {
			if stop {
				fmt.Println("Горутина с условием остановлена")
				return
			}

			fmt.Println("Горутина с условием работает...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
	stop = true
	time.Sleep(time.Second)

	// ================================================
	//           Выход через канал уведомления
	// ================================================
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("Горутина с каналом уведомления остановлена")
				return
			default:
				fmt.Println("Горутина с каналом уведомления работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(time.Second)
	close(ch)
	time.Sleep(time.Second)

	// ================================================
	//             Выход через контекст
	// ================================================
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина с контекстом остановлена")
				return
			default:
				fmt.Println("Горутина с контекстом работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)

	// ================================================
	//    Прекращение работы через runtime.Goexit()
	// ================================================
	go func() {
		fmt.Println("Горутина с Goexit работает...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Горутина с Goexit остановлена")
		runtime.Goexit()
	}()
	time.Sleep(time.Second)

	// ================================================
	//        		Выход через таймер
	// ================================================
	timer := time.After(2 * time.Second)
	go func() {
		for {
			select {
			case <-timer:
				fmt.Println("Горутина с таймером остановлена")
				return
			default:
				fmt.Println("Горутина с таймером работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)

	fmt.Println("Главная горутина main завершилась")
}
