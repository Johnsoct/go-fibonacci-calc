package main

import "fmt"

// fibonacci is a function that returns a function that returns an int
// F(n) = F(n-1) + F(n-2)
// 0 && the first 1 have no preceding numbers
func fibonacci() func(int) int {
	sequence := make([]int, 0)

	return func(index int) int {
		// I'm not understanding something available to me where I can't implicitly handle 0 and 1...
		if index == 0 {
			sequence = append(sequence, 0)
			return 0
		}

		if index == 1 {
			sequence = append(sequence, 1)
			return 1
		}
		
		firstPreceding := sequence[index-1]
		secondPreceding := sequence[index-2]
		next := firstPreceding + secondPreceding
		sequence = append(sequence, next)

		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
