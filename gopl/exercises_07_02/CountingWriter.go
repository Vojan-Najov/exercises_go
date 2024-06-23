// Exercise 7.2:
// Write a function CountingWriter with the signature below that, given an io.Writer,
// returns a new Writer that wraps the original, and a pointer to an int64 variable that
// at any moment contains the number of bytes written to the new Writer.
// func CountingWriter(w io.Writer) (io.Writer, *int64)

package main

import (
	"bytes"
	"fmt"
	"io"
)

type CountWriter struct {
	w     io.Writer
	bytes int64
}

func (cw *CountWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.bytes += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{w: w, bytes: 0}
	return &cw, &cw.bytes
}

func main() {
	var buf bytes.Buffer
	cw, nptr := CountingWriter(&buf)
	fmt.Fprintf(cw, "Hello %s", "world")
	fmt.Println(buf.String(), *nptr)
	fmt.Fprintf(cw, "!!!")
	fmt.Println(buf.String(), *nptr)
}
