package main

func GetFactorial(n int) (int) {
	if n < 0 {
		return 0
	}

	if n <= 1 {
		return 1
	}

	return n * GetFactorial(n - 1)
}