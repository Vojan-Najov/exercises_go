/*
Любимые студенты

Мы решили посчитать аналитику и посмотреть, а сколько же с нами учится
каждый студент курса по Go - то есть найти кол-во дней, которое он
(студент) провел в курсе с момента поступления и до 1 октября 2023 года.

Напишите функцию
processJSON(jsonData []byte) error,
которая должна принимать данные о студентах в формате JSON, разбирать их и
выводить искомое число дней.

Вывод должен быть в формате имяСтудента : количество дней
Формат ввода
[
    {
        "name": "Анна",
        "admission_date": "2023-09-29T00:00:00Z"
    },
    {
        "name": "Иван",
        "admission_date": "2023-09-28T00:00:00Z"
    }
]
Формат вывода
Анна: 2
Иван: 3
Примечания
type Student struct {
    Name         string    `json:"name"`
    AdmissionDate time.Time `json:"admission_date"`
    DaysOnCourse int       `json:"-"`
}
*/

package main

import (
  "fmt"
  "time"
  "encoding/json"
)

type Student struct {
    Name          string    `json:"name"`
    AdmissionDate time.Time `json:"admission_date"`
    DaysOnCourse  int       `json:"-"`
}

func processJSON(jsonData []byte) error {
  var students []Student

  if err := json.Unmarshal(jsonData, &students); err != nil {
    return err
  }  

  end := time.Date(2023, time.October, 1, 0, 0, 0, 0, time.UTC)
  for _, student := range students {
    start := student.AdmissionDate
    length := end.Sub(start).Hours() / 24.0
    fmt.Printf("%s: %d\n", student.Name, int(length))
  }

  return nil
}

func main() {
  input := []byte(`[
    {
        "name": "Анна",
        "admission_date": "2023-09-29T00:00:00Z"
    },
    {
        "name": "Иван",
        "admission_date": "2023-09-28T00:00:00Z"
    }
  ]`)

  if err := processJSON(input); err != nil {
    fmt.Println("Error: ", err)
  }
}



