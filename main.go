package readcompare

import (
	"bufio"
	"bytes"
	"io"
)

func Comparer1(r1, r2 io.Reader) (bool, error) {
	br1 := bufio.NewReader(r1)
	br2 := bufio.NewReader(r2)
	for {
		c1, err1 := br1.ReadByte()
		c2, err2 := br2.ReadByte()
		if err1 != nil {
			if err1 == io.EOF {
				if err2 == io.EOF {
					return true, nil
				}
				return false, err2
			}
			return false, err1
		}
		if err2 != nil {
			return false, err2
		}
		if c1 != c2 {
			return false, nil
		}
	}
}

func Comparer2(r1, r2 io.Reader) (bool, error) {
	const UNIT = bufio.MaxScanTokenSize // 64*1024
	var buffer1 [UNIT]byte
	var buffer2 [UNIT]byte

	for {
		n1, err1 := io.ReadFull(r1, buffer1[:])
		n2, err2 := io.ReadFull(r2, buffer2[:])

		if err1 != nil {
			if err1 == io.EOF || err1 == io.ErrUnexpectedEOF {
				if err2 == io.EOF || err2 == io.ErrUnexpectedEOF {
					return n1 == n2 && bytes.Equal(buffer1[:n1], buffer2[:n2]), nil
				}
				return false, err2
			}
			return false, err1
		}
		if err2 != nil {
			return false, err2
		}
		if n1 != n2 || !bytes.Equal(buffer1[:n1], buffer2[:n2]) {
			return false, nil
		}
	}
}
