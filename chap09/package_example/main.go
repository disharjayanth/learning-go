package main

import (
	"fmt"

	print "github.com/disharjayanth/learningGo/chap09/package_example/formatter"
	"github.com/disharjayanth/learningGo/chap09/package_example/math"
)

func main() {
	var num int
	fmt.Println("Enter the number to be doubled:")
	fmt.Scanf("%d", &num)
	doubledNum := math.Double(num)
	stringFormat := print.Format(doubledNum)
	fmt.Println(stringFormat)
}
