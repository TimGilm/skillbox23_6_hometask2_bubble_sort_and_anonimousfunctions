/*
Задание 2. Анонимные функции
Что нужно сделать Напишите анонимную функцию, которая на
вход получает массив типа integer, сортирует его пузырьком
и переворачивает (либо сразу сортирует в обратном порядке,
как посчитаете нужным).
*/
// первый вариант, сперва сортировка пузырьком, затем реверсирование
package main

import "fmt"

const size = 10

func main() {
	a := [size]int{8, 9, 3, 5, 2, 6, 1, 7, 4, 0}
	fmt.Println(a)
	f := func(a [size]int) [size]int {
		for i := size; i > 0; i-- {
			for j := 1; j < i; j++ {
				if a[j-1] < a[j] {
					a[j-1], a[j] = a[j], a[j-1]
				}
			}
		}
		return a
	}
	fmt.Println(f(a))
}
