package main

import "fmt"

func main() {
	i, j := 100, 200
	p := &i
	fmt.Println(*p)
	*p = 101
	fmt.Println(i)

	p = &j
	*p = *p / 10
	fmt.Println(j)
}
