package main

import (
    "fmt"
)

func countOccurrences(nums []int) map[int]int {
    occurrences := make(map[int]int)

    for _, num := range nums {
        occurrences[num]++
    }

    return occurrences
}

func main() {
    numbers := []int{1, 2, 3, 2, 1, 3, 4, 5, 4, 1}
    result := countOccurrences(numbers)

    fmt.Println("Результат подсчёта повторений:")
    for num, count := range result {
        fmt.Printf("%d встречается %d раз(а)\n", num, count)
    }
}
