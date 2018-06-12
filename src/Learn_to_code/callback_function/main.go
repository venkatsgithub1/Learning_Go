package main

// earlier an external function.
func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
}
