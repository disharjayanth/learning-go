package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func buildGZipReader(filename string) (*gzip.Reader, func(), error) {
	file, err := os.Open(filename)
	fmt.Println("file:", file)
	if err != nil {
		return nil, nil, err
	}
	gzipReader, err := gzip.NewReader(file)
	fmt.Println("gzipReader:", gzipReader, "Error:", err)
	if err != nil {
		return nil, nil, err
	}
	return gzipReader, func() {
		file.Close()
		gzipReader.Close()
	}, nil
}

func countLetters(r io.Reader) (map[string]int, error) {
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
	gzipReader, closer, err := buildGZipReader("someTxt.zip")
	checkError(err)
	defer closer()
	counts, err := countLetters(gzipReader)
	checkError(err)
	fmt.Println("Counts:", counts)
}
