/*
Чтение JSON в Go

Для чтения JSON в Go нужна функция Unmarshal() из пакета encoding/json.
Эта функция принимает на вход байтовый массив, который содержит JSON-данные,
и указатель на переменную, в которую будут записаны декодированные данные.

В этом примере мы создали структуру Person с полями Name и Age.
Затем определили строку jsonStr, которая содержит данные в JSON.
Мы распаковали данные из строки JSON с помощью Unmarshal() и записали их в переменную person.

Обратите внимание: мы передаём переменную person по указателю (символ &).
Сигнатура функции Unmarshal при этом возвращает только ошибку.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// `json:"name"` — тег поля. Без него в JSON будет ключ "Name" вместо "name"
	Name string `json:"name"` 
	Age int `json:"age"`
	// поле с тегом `json:"-"` при кодировании json игнорируется
	Gender string `json:"-"`
	// неэкспортируемые поля так же игнорируются
	privateNotes string
}

func main() {
	jsonStr := `{"name": "John", "age": 30, "Gender": "male"}`
	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)
}
