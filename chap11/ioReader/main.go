package main

import (
	"fmt"
	"io"
	"strings"
)

func countLetter(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func main() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetter(sr)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Count letters:", counts)
}
