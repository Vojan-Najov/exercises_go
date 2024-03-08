/*
CSV

CSV (от англ. Comma-Separated Values — значения, разделённые запятыми) — текстовый
формат, предназначенный для представления табличных данных. Строка таблицы
соответствует строке текста, которая содержит одно или несколько полей, разделенных
запятыми.

Напишите функцию SumUp(filepath, colname string) (int, error), которая читает файл
формата csv и суммирует значения из колонки colname. Верните полученную сумму если
нет ошибок, иначе 0 и ошибку.

Примечания
первая строка файла - имена колонок.

*/

package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func SumUp(filepath, colname string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	r := csv.NewReader(file)

	names, err := r.Read()
	if err == io.EOF {
		return 0, errors.New("empty file")
	}
	if err != nil {
		return 0, err
	}

	idx := slices.Index(names, colname)
	if idx == -1 {
		return 0, errors.New("unknown colname")
	}

	sum := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		num, err := strconv.Atoi(record[idx])
		if err != nil {
			return 0, err
		}
		sum += num
	}

	return sum, nil
}

func main() {
	sum, err := SumUp("tmp", "first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sum is %d\n", sum)
	}

	sum, err = SumUp("tmp", "second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sum is %d\n", sum)
	}

	sum, err = SumUp("tmp", "third")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sum is %d\n", sum)
	}

	sum, err = SumUp("tmp", "noexist")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sum is %d\n", sum)
	}

	sum, err = SumUp("noexist", "noexist")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sum is %d\n", sum)
	}
}
