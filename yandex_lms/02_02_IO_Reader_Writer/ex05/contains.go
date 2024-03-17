/*
Reader_Writer_5

Напишите функцию Contains(r io.Reader, seq []byte) (bool, error) которая должна найти
первое вхождение байт seq в данных, доступных через Reader r. Если последовательность
найдена - верните true, nil, иначе false, nil.
В случае возникновения ошибки - false и возникшую ошибку.
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	var err error
	var nread int
	var offset int
	buf := make([]byte, 0, 2*len(seq))
	data := make([]byte, len(seq))

	for err != io.EOF {
		nread, err = r.Read(data)
		if err != nil && err != io.EOF {
			return false, err
		}
		buf = append(buf[offset:], data[:nread]...)

		if bytes.Contains(buf, seq) {
			return true, nil
		}
		if offset == 0 {
			offset = 1
		} else {
			offset = len(seq)
		}
	}
	return false, nil
}

func main() {
	r := strings.NewReader("abcdefg")
	fmt.Println(Contains(r, []byte("de")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(r, []byte("123")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(r, []byte("q")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(r, []byte("a")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(r, []byte("b")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(r, []byte("g")))
}
