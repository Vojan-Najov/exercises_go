package main

import (
	"context"
	"io"
	"slices"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	var buf []byte

	smallBuf := make([]byte, 1)

	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
			_, err := r.Read(smallBuf)
			if err != nil {
				if err == io.EOF {
					return false, nil
				}
				return false, err
			}

			buf = append(buf, smallBuf...)

			if len(buf) > len(seq) {
				buf = buf[1:]
			}

			if slices.Equal(buf, seq) {
				return true, nil
			}
		}
	}
}
