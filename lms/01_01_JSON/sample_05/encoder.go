/*

Кодирование JSON с помощью Writer

Вот пример кода, который записывает JSON-данные в буфер с помощью
json.Кодировщик:
В этом примере мы создаём слайс структур Student с данными о школьниках,
затем — буфер для записи JSON-данных и Encoder для записи этих данных в буфер.
Затем мы записываем их в буфер по методу Encode() Encoder и выводим результат
на экран в виде JSON-данных о студентах.

*/

package main

import (
  "fmt"
  "encoding/json"
  "bytes"
)

type Student struct {
  Name string `json:"name"`
  Age int     `json:"age"`
  Grade int   `json:"grade"`
}

func main() {
  // Создаём слайс структур Student с данными о школьниках
  students := []Student{
    {Name: "Alice", Age: 12, Grade: 7},
    {Name: "Bob", Age: 13, Grade: 8},
    {Name: "Charlie", Age: 14, Grade: 9},
  }

  // Создаём буфер для записи JSON-данных
  var buf bytes.Buffer
  
  // Создаём Encoder для записи JSON-данных в буфер
  encoder := json.NewEncoder(&buf)

  // Записываем JSON-данные в буфер с помощью метода Encode() Encoder
  err := encoder.Encode(students)
  if err != nil {
    fmt.Println("Ошибка при запими JSON-данных: ", err)
    return
  }

  // Выводим результат на экран
  fmt.Println("Данные о студентах:")
  fmt.Println(buf.String());
}

