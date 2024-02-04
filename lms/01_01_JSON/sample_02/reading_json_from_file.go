/*
Запись JSON в Go

Для записи данных в JSON в Go нужна функция Marshal() из пакета encoding/json.
Эта функция принимает на вход переменную с данными и возвращает байтовый массив с данными в этом формате.
Давайте так же попробуем прочитать данные из файла:
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Создаем структуру для хранения данных из JSON-файла
type Person struct {
	Name string `json:"name"`   // Поле Name для хранения имени
	Email string `json:"email"` // Поле Email для хранения адреса электронной почты
}

func main() {
	file, err := os.Open("data.json") // Открываем файл "data.json"
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file) // Считываем содержимое файла

	var person Person                        // Создаем переменную для хранения данных из JSON
	err = json.Unmarshal(byteValue, &person) // Декодируем JSON в структуру Person
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	fmt.Println("Имя:", person.Name)    // Выводим имя из структуры Person
	fmt.Println("Email:", person.Email) // Выводим адрес электронной почты из структуры Person
}
