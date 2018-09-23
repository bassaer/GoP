package main

import "fmt"

func show(s []int) {
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
}

func main() {
	src := []string{
		"hoge",
		"foo",
		"bar",
		"piyo",
	}
	a := src[0:2]
	b := src[1:3]
	fmt.Println(a, b)

	b[0] = "CHANGE"
	fmt.Println(a, b)
	fmt.Println(src)

	piyo := []bool{true, false, false, true, false}
	for i, v := range piyo {
		fmt.Printf("%d -> %v\n", i, v)
	}

	foo := make([]int, 0, 5)
	fmt.Println(foo)
	show(foo)
	foo = append(foo, 0, 2, 2)
	fmt.Println(foo)
	foo = foo[1:]
	show(foo)
	fmt.Println(foo)
}
