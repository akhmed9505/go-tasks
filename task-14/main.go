package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в runtime способна определить тип переменной,
переданной в неё (на вход подаётся interface{}).
Типы, которые нужно распознавать: int, string, bool, chan (канал).
Подсказка: оператор типа switch v.(type) поможет в решении.
*/

func GetTypeName(i interface{}) {
	switch val := i.(type) {
	case int:
		fmt.Println("int:", val)
	case string:
		fmt.Println("string:", val)
	case bool:
		fmt.Println("bool:", val)
	default:
		t := reflect.TypeOf(i)
		if t.Kind() == reflect.Chan {
			fmt.Println("chan:", t.String())
		} else {
			fmt.Println("unknown type:", t)
		}
	}
}

func main() {
	GetTypeName(1)
	GetTypeName("WB")
	GetTypeName(true)
	GetTypeName(make(chan int))
	GetTypeName(make(chan bool))
	GetTypeName(1.0)
}
