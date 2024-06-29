// Exercise 7.5:
// The LimitReader function in the io package accepts an io.Reader r and a number of
// bytes n, and returns another Reader that reads from r but reports an end-of-file
// condition after n bytes. Implement it.
// func LimitReader(r io.Reader, n int64) io.Reader

package limit_reader

import (
	"io"
)

type LimitReaderT struct {
	r   io.Reader
	n   int64
	cur int64
}

func (r *LimitReaderT) Read(p []byte) (int, error) {
	if r.cur == r.n {
		return 0, io.EOF
	}
	size := int64(len(p))
	if size > r.n-r.cur {
		size = r.n - r.cur
	}
	r.cur += size
	return r.r.Read(p[:size])
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitReaderT{r: r, n: n, cur: 0}
}
