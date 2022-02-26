package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibnacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	sprintf := fmt.Sprintf("%d\n", next)
	read, err := strings.NewReader(sprintf).Read(p)
	return read, err
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibnacci()
	//for i := 0; i < 20; i++ {
	//	fmt.Println(f())
	//}
	printFileContents(f)
}
