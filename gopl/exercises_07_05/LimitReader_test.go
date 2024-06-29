package limit_reader

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("0123456789")
	lr := LimitReader(r, 5)
	p := make([]byte, 4)

	n, err := lr.Read(p)
	if n != 4 || err != nil {
		t.Errorf(
			"Expected n: 4, err: nil; got n: %d, err: %v",
			n,
			err,
		)
	}

	n, err = lr.Read(p)
	if n != 1 || err != nil {
		t.Errorf(
			"Expected n: 4, err: nil; got n: %d, err: %v",
			n,
			err,
		)
	}

	n, err = lr.Read(p)
	if n != 0 || err != io.EOF {
		t.Errorf(
			"Expected n: 4, err: nil; got n: %d, err: %v",
			n,
			err,
		)
	}
}
