package main

import "fmt"

func end() {
	fmt.Println("end")
}

func main() {
	defer end()
	fmt.Println("start")
	fmt.Println("running..")
}
