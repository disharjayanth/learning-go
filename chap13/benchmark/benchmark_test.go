package benchmark

import (
	"fmt"
	"testing"
)

var blackhole int

func TestFileLen(t *testing.T) {
	charCount, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if charCount != 12 {
		t.Errorf("Expected %d, got %d\n", 12, charCount)
	}
}

func BenchmarkFileLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		charCount, err := FileLen("testdata/data.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = charCount
	}
}

func BenchmarkFileLen2(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countChar, err := FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = countChar
			}
		})
	}
}
