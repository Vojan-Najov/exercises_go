/*
С новым учебным годом!

1 сентября каждого учебного года во всех базах данных школьников происходит
великий пересчёт. Напишите функцию
modifyJSON(jsonData []byte) ([]byte, error),
которая принимает данные в формате JSON, добавляет 1 год к полю grade и
возвращает обновлённые данные также в формате JSON.

Формат ввода
[
    {
        "name": "Oleg",
        "grade": 9
    },
    {
        "name": "Katya",
        "grade": 10
    }
]

Формат вывода
[
    {
        "name": "Oleg",
        "grade": 10
    },
    {
        "name": "Katya",
        "grade": 11
    }
]

Примечания

Структура ученика

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}
*/

package main

import (
  "encoding/json"
  "fmt"
)

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
  var students []Student

  if err := json.Unmarshal(jsonData, &students); err != nil {
    return nil, err
  }

  for i := 0; i < len(students); i++ {
    students[i].Grade++
  }

  jsonBytes, err :=  json.Marshal(students)
  if err != nil {
    return nil, err
  }

  return jsonBytes, nil
}

func main() {
  students := []Student{
    {Name: "A", Grade: 10},
    {Name: "B", Grade: 12},
  }
  
  jsonBytes, err := json.Marshal(students)
  if err != nil {
    fmt.Println("Marshal error: ", err)
    return
  }
  fmt.Println(string(jsonBytes))

  jsonBytes, err = modifyJSON([]byte(jsonBytes))
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(jsonBytes))
}

