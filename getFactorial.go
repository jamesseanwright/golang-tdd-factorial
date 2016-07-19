package main

func GetFactorial(n int) (int) {
	if n < 0 {
		return 0
	}

	if n <= 1 {
		return 1
	}

	factorial := n

	for i := n - 1; i > 0; i-- {
		factorial *= i
	}

	return factorial
}