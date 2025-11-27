package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/

type human struct {
	name string
	age  int
	work string
}

func (h *human) Info() string {
	return fmt.Sprintf("My name is %s, I'm %v. I work as a %s!", h.name, h.age, h.work)
}

type Action struct {
	human
}

func main() {
	h := &human{
		name: "John",
		age:  30,
		work: "programmer",
	}

	a := Action{*h}
	fmt.Println(a.Info())
}
