package main

import "fmt"

func quickSort(slice []int) []int {
    // Если длина списка меньше 2, он уже отсортирован
    if len(slice) < 2 {
       return slice
    }

    // Выбираем опорный элемент (обычно это первый или последний элемент)
    pivot := slice[0]

    // Создаём пустые списки для элементов меньше и больше опорного элемента
    var less []int
    var greater []int

    // Итерируемся по списку и добавляем элементы в соответствующие списки
    for _, element := range slice[1:] {
       if element <= pivot {
          less = append(less, element)
       } else {
          greater = append(greater, element)
       }
    }

    // Рекурсивно вызываем быструю сортировку для каждого из списков
    // и объединяем результаты вместе с опорным элементом
    return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func main() {
    // Создаём неотсортированный срез целых чисел
    s := []int{5, 2, 4, 6, 1, 3}

    // Печатаем неотсортированный срез
    fmt.Println("Неотсортированный:", s)

    // Сортируем срез с помощью быстрой сортировки
    sorted := quickSort(s)

    // Печатаем отсортированный срез
    fmt.Println("Отсортированный:", sorted)
}
