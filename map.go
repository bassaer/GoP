package main

import "fmt"

func show(m map[string]bool) {
	for k, v := range m {
		fmt.Printf("key: %s, value: %v\n", k, v)
	}
}

func exists(m map[string]bool, k string) {
	v, ok := m[k]
	if ok {
		fmt.Printf("exists! -> %v\n", v)
	} else {
		fmt.Println("not exists")
	}
}

func main() {
	m := make(map[string]bool)
	m["hoge"] = false
	m["piyo"] = true
	m["foo"] = false
	show(m)
	exists(m, "piyo")
	delete(m, "piyo")
	show(m)
	exists(m, "piyo")
}
