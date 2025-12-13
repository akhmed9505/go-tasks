package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
По завершению программы структура должна выводить итоговое значение счётчика.
Подсказка: вам понадобится механизм синхронизации,
например, sync.Mutex или sync/Atomic для безопасного инкремента.
*/

// ============ Mutex ============
type mutexCounter struct {
	val int
	mu  sync.Mutex
}

func (mc *mutexCounter) Inc(wg *sync.WaitGroup) {
	defer wg.Done()
	mc.mu.Lock()
	mc.val++
	mc.mu.Unlock()
}

func (mc *mutexCounter) Value() int {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	return mc.val
}

// ============ Atomic ============
type atomicCounter struct {
	val int64
}

func (ac *atomicCounter) Inc(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&ac.val, 1)
}

func (ac *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&ac.val)
}

// ============ Main ============
func main() {
	var wg sync.WaitGroup

	mu := mutexCounter{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go mu.Inc(&wg)
	}
	wg.Wait()
	fmt.Println("Mutex result:", mu.Value())

	ac := atomicCounter{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go ac.Inc(&wg)
	}
	wg.Wait()
	fmt.Println("Atomic result:", ac.Value())
}
