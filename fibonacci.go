package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	a := []int{0, 1}
	return func() int {
		switch i {
		case 0:
			i++
			return 0
		case 1:
			i++
			return 1
		default:
			n := a[i-1] + a[i-2]
			i++
			a = append(a, n)
			return n
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
