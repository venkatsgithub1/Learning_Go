package main

import "fmt"

func decrement() func() int {
	x := 0
	return func() int {
		x--
		return x
	}
}

func evenNumberGenerator() func() uint {
	i := uint(0)

	return func() (retVar uint) {
		retVar = i
		i += 2
		return retVar
	}
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 1))
	x := 1
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	decrease := decrement()

	fmt.Println(decrease())

	fmt.Println(decrease())

	generateAnEvenNumber := evenNumberGenerator()

	fmt.Println(generateAnEvenNumber())

	fmt.Println(generateAnEvenNumber())

	x1 := 0
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(x1 + n)
	})

	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)
}
