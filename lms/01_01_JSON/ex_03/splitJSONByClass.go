/*
Равнение на флаг!

На линейке ученикам нужно сгруппироваться по классам.
Проведём линейку для базы данных.
Напишите функцию
splitJSONByClass(jsonData []byte) (map[string][]byte, error),
которая принимает данные в формате JSON и возвращает мапу, в которой ключи —
классы, а значения — данные в формате JSON, которые к этому классу
относятся.

Примечания
Например: Входные данные
[
    {
        "name": "Oleg",
        "class": "9B"
    },
    {
        "name": "Ivan",
        "class": "9A"
    },
    {
        "name": "Maria",
        "class": "9B"
    },
    {
        "name": "John",
        "class": "9A"
    }
]
Выходные данные должны быть в виде map:
map[string][]byte{
  "9A": []byte(`[
                  {"class":"9A","name":"Ivan"},
                  {"class":"9A","name":"John"}
                ]`),
  "9B": []byte(`[
                  {"class":"9B","name":"Oleg"},
                  {"class":"9B","name":"Maria"}
                ]`),
}
*/

package main

import (
  "encoding/json"
  "fmt"
)

type Student struct {
  Class string `json:"class"`
  Name string `json:"name"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
  var students []Student

  if err := json.Unmarshal(jsonData, &students); err != nil {
    return nil, err
  }

  var m map[string][]byte
  m = make(map[string][]byte)
  for _, student := range students {
    data, err := json.Marshal(student)
    if err != nil {
      return nil, err
    }

    value, ok := m[student.Class]
    if ok {
      m[student.Class] = append(value, data...)
    } else {
      m[student.Class] = data
    }
  }

  return m, nil
}

func main() {
  input := []byte(`[
    {
        "name": "Oleg",
        "class": "9B"
    },
    {
        "name": "Ivan",
        "class": "9A"
    },
    {
        "name": "Maria",
        "class": "9B"
    },
    {
        "name": "John",
        "class": "9A"
    }
  ]`)
  res := map[string][]byte{
    "9A": []byte(`[
                    {"class":"9A","name":"Ivan"},
                    {"class":"9A","name":"John"}
                  ]`),
    "9B": []byte(`[
                    {"class":"9B","name":"Oleg"},
                    {"class":"9B","name":"Maria"}
                  ]`),
    }

  m, err := splitJSONByClass(input)
  if err != nil {
    fmt.Println(err)
    return
  }

  for k, v := range res {
    fmt.Println(k, string(v))
  }
  fmt.Println()
  for k, v := range m {
    fmt.Println(k, string(v))
  }
}
