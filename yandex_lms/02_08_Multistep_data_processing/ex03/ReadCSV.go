/*
 * 3. Чтение CSV
 *
 * Напишите функцию ReadCSV(file string) (<-chan []string, error) для чтения
 * csv-файлов. В выходной канал должны передаваться строки из файла. Если
 * возникла ошибка, верните ее описание.
 */

package main

import (
	"encoding/csv"
	"os"
)

func ReadCSV(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	out := make(chan []string)
	go func() {
		defer close(out)
		for _, record := range records {
			out <- record
		}
	}()

	return out, nil
}
