package examples

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func GoRecursion() {
	fmt.Println(fact(7))
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			// fmt.Println("return ", n)
			return n
		}
		rs := fib(n-1) + fib(n-2)
		// fmt.Println("return ", rs)
		return rs
	}
	fmt.Println(fib(7))
}
