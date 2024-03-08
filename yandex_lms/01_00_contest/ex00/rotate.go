/*
Сдвиг

Напишите функцию Rotate(data []int, pos int) []int, которая осуществляет циклический
сдвиг элементов слайса чисел на заданное количество позиций.
Пример: если data = [1,2,3,4,5,6,7] и pos = 3, то функция должна вернуть [5,6,7,1,2,3,4]

Примечания
pos может быть отрицательным числом, а также больше длины слайса.
*/

package main

import "fmt"

func reverse(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func Rotate(data []int, pos int) []int {
	rotate := make([]int, len(data))
	copy(rotate, data)

	pos %= len(data)
	if pos < 0 {
		pos += len(data)
	}
	pos = len(data) - pos

	reverse(rotate[:pos])
	reverse(rotate[pos:])
	reverse(rotate)

	return rotate
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	for i := -14; i < 15; i++ {
		rotate := Rotate(data, i)
		fmt.Printf("%3d: %v\n", i, rotate)
	}
}
