/*
Колобок

По амбарам помела, по сусекам поскребла — базу данных получила.
Напишите функцию mergeJSONData(jsonDataList ...[]byte) ([]byte, error),
которая принимает несколько экземпляров данных в формате JSON, объединяет их
в один экземпляр и возвращает его.
Примечания:
Например: В функцию передаются два JSON:
[
    {
        "name": "Oleg",
        "class": "9B"
    },
    {
        "name": "Ivan",
        "class": "9A"
    }
]
и
[
    {
        "name": "Maria",
        "class": "9B"
    },
    {
        "name": "John",
        "class": "9A"
    }
]
На выходе нужно получить:
[
    {
        "class": "9B",
        "name": "Oleg"
    },
    {
        "class": "9A",
        "name": "Ivan"
    },
    {
        "class": "9B",
        "name": "Maria"
    },
    {
        "class": "9A",
        "name": "John"
    }
]
*/

package main

import (
  "encoding/json"
  "fmt"
)

type Student struct {
  Class string `json:"class"`
  Name string  `json:"name"`
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
  var allStudents []Student
  for _, jsonData := range jsonDataList {
    var students []Student
    if err := json.Unmarshal(jsonData, &students); err != nil {
      return nil, err
    }
    for _, student := range students {
      allStudents = append(allStudents, student)
    }
  }

  jsonData, err := json.Marshal(allStudents)
  if err != nil {
    return nil, err
  }

  return jsonData, nil
}

func main() {
  inputJSON1 := []byte(`[
    {
      "name": "Oleg",
      "class": "9B"
    },
    {
      "name": "Ivan",
      "class": "9A"
    }
  ]`)
  inputJSON2 := []byte(`[
    {
      "name": "Maria",
      "class": "9B"
    },
    {
      "name": "John",
      "class": "9A"
    }
  ]`)

  jsonData, err := mergeJSONData(inputJSON1, inputJSON2)
  if err != nil {
    fmt.Println(err);
  }
  fmt.Println(string(jsonData))
}
