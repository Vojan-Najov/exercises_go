/*

Маршаллинг

Для записи данных в формате JSON в Go необходимо использовать функцию Marshal() из пакета encoding/json.
Эта функция принимает на вход переменную, содержащую данные, которые нужно записать в формате JSON,
и возвращает байтовый массив с данными в формате JSON.

*/

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	person := Person{Name: "John", Age: 30}
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}
