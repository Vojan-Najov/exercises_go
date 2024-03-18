/*
Поиск последовательности

Напишите функцию Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error)
которая должна найти первое вхождение байт seq в данных, доступных через Reader r.
Если последовательность найдена - верните true, nil, иначе false, nil. В случае
возникновения ошибки - false и возникшую ошибку. В случае отмены контекста - функция
должна вернуть false и ошибку - причину отмены контекста.
*/

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	var err error
	var nread int
	var offset int
	buf := make([]byte, 0, 2*len(seq))
	data := make([]byte, len(seq))

	for err != io.EOF {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
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
	}
	return false, nil
}

func main() {
	r := strings.NewReader("abcdefg")
	fmt.Println(Contains(context.TODO(), r, []byte("de")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(context.TODO(), r, []byte("123")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(context.TODO(), r, []byte("q")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(context.TODO(), r, []byte("a")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(context.TODO(), r, []byte("b")))
	r = strings.NewReader("abcdefg")
	fmt.Println(Contains(context.TODO(), r, []byte("g")))
}
