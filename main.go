/*
Задание 2. Анонимные функции
Что нужно сделать Напишите анонимную функцию, которая на
вход получает массив типа integer, сортирует его пузырьком
и переворачивает (либо сразу сортирует в обратном порядке,
как посчитаете нужным).
*/
package main

import "fmt"

const size = 10

func main() {
	a := [size]int{8, 9, 3, 5, 2, 6, 1, 7, 4, 0}
	fmt.Println(a)
	fmt.Println(bubbleSortAndInvert(a))
}

func bubbleSortAndInvert(a [size]int) [size]int {
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	return a
}
