Benchmark comparing the contents two io.Readers
===============================================

main.go
-------

```main.go
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
```

go test -bench . -benchmem
--------------------------

```go test -bench . -benchmem|
goos: windows
goarch: amd64
pkg: github.com/hymkor/study-go-read-comparer
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
BenchmarkComparer1-4   	      24	  48393312 ns/op	    9536 B/op	       8 allocs/op
BenchmarkComparer2-4   	     392	   3026203 ns/op	  132416 B/op	       8 allocs/op
PASS
ok  	github.com/hymkor/study-go-read-comparer	2.891s
```

The sample data `sample1.bin` and `sample2.bin` used in [main\_test.go](main_test.go) are not commited in this repository, so please prepare them yourself.
