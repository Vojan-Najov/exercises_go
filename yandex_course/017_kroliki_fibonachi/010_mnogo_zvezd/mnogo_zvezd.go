package main

import "fmt"

func main() {
    var count int
    
    if _, err := fmt.Scanln(&count); err != nil || count < 0 {
        fmt.Println("Некорректный ввод")
        return
    }
    
    switch {
        case count < 10:
        fmt.Println("Число меньше 10")
        case count < 100:
        fmt.Println("Число меньше 100")
        case count < 1000:
        fmt.Println("Число меньше 1000")
        default:
        fmt.Println("Число больше или равно 1000")
    }
}
        
