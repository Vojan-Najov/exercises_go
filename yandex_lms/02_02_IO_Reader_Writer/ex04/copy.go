/*
Reader_Writer_4

Напишите функцию Copy(r io.Reader, w io.Writer, n uint) error, которая копирует n
байт из r в w.
Если количество байт, доступных для чтения меньше n - функция должна копировать все
данные. В случае ошибки - верните её.
*/

package main

import "io"

func Copy(r io.Reader, w io.Writer, n uint) error {
	data := make([]byte, n)
	nread, err := r.Read(data)
	if err != nil && err != io.EOF {
		return err
	}
	var offset int
	for nread > 0 {
		nwrite, err := w.Write(data[offset:nread])
		if err != nil {
			return err
		}
		nread -= nwrite
		offset += nwrite
	}
	return nil
}
