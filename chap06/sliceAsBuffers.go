package main

import (
	"fmt"
	"os"
)

func process(data []byte) {
	fmt.Println(string(data))
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error in opening file:", err)
	}
	defer file.Close()

	data := make([]byte, 900)
	i := 0
	fmt.Println("Read content from file:")
	for {
		n, err := file.Read(data)
		i++
		if err != nil {
			fmt.Println("Cannot read from file:", err)
			break
		}
		if n == 0 {
			fmt.Println("End of file....")
			break
		}
		process(data[:n])
		// fmt.Println("Number of read loops:", i)
	}
}
