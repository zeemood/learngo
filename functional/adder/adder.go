package main

import "fmt"

func adder() func(val int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	b := adder2(0)
	for i := 1; i <= 10; i++ {
		fmt.Printf("0+...+%d=%d\n", i, a(i))
		var s int
		s, b = b(i)
		fmt.Printf("0+...+%d=%d\n", i, s)
	}

}
