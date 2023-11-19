package readcompare

import (
	"io"
	"os"
	"testing"
)

func tryCompare(b *testing.B, comparer func(r1, r2 io.Reader) (bool, error)) bool {
	b.Helper()
	fd1, err := os.Open("sample1.bin")
	if err != nil {
		b.Fatal(err.Error())
	}
	defer fd1.Close()

	fd2, err := os.Open("sample2.bin")
	if err != nil {
		b.Fatal(err.Error())
	}
	defer fd2.Close()

	value, err := comparer(fd1, fd2)
	if err != nil {
		b.Fatal(err.Error())
	}
	return value
}

func benchmarkComparer(b *testing.B, comparer func(r1, r2 io.Reader) (bool, error)) {
	b.Helper()
	for i := 0; i < b.N; i++ {
		if !tryCompare(b, comparer) {
			b.Fatal("not equal")
		}
	}
}

func BenchmarkComparer1(b *testing.B) {
	benchmarkComparer(b, Comparer1)
}

func BenchmarkComparer2(b *testing.B) {
	benchmarkComparer(b, Comparer2)
}
