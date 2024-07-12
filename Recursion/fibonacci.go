package main

import "fmt"

func main() {
	result := tail(5, 0, 0)
	fmt.Println(result)
}

func tail(n int, a int, b int) int {

	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	fib1 := tail(n-1, a, b)
	fib2 := tail(n-2, a, b)

	result := fib1 + fib2

	return result

}
