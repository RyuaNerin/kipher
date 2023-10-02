package lsh256

import (
	"testing"
)

const shortWriteSize = 16 * 1024

func Test_ShortWrite_Go(t *testing.T) {
	testSize(
		t,
		func(t *testing.T, size int) {
			h := newContextGo(size)

			buf := make([]byte, shortWriteSize)
			for i := 1; i < shortWriteSize; i++ {
				rnd.Read(buf[:i])
				n, err := h.Write(buf[:i])
				if err != nil {
					t.Error(err)
				}
				if n != i {
					t.Fail()
				}

				rnd.Read(buf[:1])
				switch buf[0] % 5 {
				case 0:
					h.Reset()
				case 1:
					h.Sum(buf[:0])
				}
			}
		},
	)
}
