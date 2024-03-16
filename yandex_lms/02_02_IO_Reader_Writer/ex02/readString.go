/*
Reader_Writer_2

Напишите функцию ReadString(r io.Reader) (string, error), которая читает данные с
помощью r и возвращает эти данные в строковом виде.
В случае возникновения ошибки функция должна вернуть пустую строку и возникшую
ошибку, иначе строку и nil.
*/

package main

import "io"

func ReadString(r io.Reader) (string, error) {
	var n int
	var err error
	var readen []byte
	data := make([]byte, 1024)
	for {
		n, err = r.Read(data)
		readen = append(readen, data[:n]...)
		if err != nil {
			break
		}
	}
	if err == io.EOF {
		err = nil
	}
	return string(readen), err
}
