package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
Реализовать безопасную для конкуренции запись данных в структуру map.
Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
Проверьте работу кода на гонки (util go run -race).
*/

type Data struct {
	mu sync.Mutex
	m  map[string]int
}

func NewData() *Data {
	return &Data{
		m: make(map[string]int),
	}
}

func (d *Data) Set(key string, value int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.m[key] = value
}

func (d *Data) Get(key string) (int, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if value, ok := d.m[key]; ok {
		return value, ok
	}
	return 0, false
}

func (d *Data) Len() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return len(d.m)
}

func main() {
	var wg sync.WaitGroup
	data := NewData()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := strconv.Itoa(id)
			data.Set(key, id*id)

			if value, ok := data.Get(key); ok {
				fmt.Printf("Записано: %s: %d\n", key, value)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Итоговое количество элементов: %d", data.Len())
}
