package main

import "fmt"

func main() {
    var t1 int
    var t2 int
    var t3 int
    _, err := fmt.Scanln(&t1, &t2, &t3)
    if err != nil {
      fmt.Println(err)
      return
    } else if (t1 < 0 || t2 < 0 || t3 < 0) {
      fmt.Println("Некорректный ввод")
      return
    }
    
    if t1 == t2 && t1 == t3 {
        fmt.Println("Все числа равны")
    } else if t1 == t2 || t1 == t3 || t2 == t3 {
        fmt.Println("Два числа равны")
    } else {
        fmt.Println("Все числа разные")
    }
}
        
