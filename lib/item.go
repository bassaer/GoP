package lib

import (
	"fmt"
)

type Item struct {
	Id   int
	Name string
}

func (i *Item) Show() {
	fmt.Printf("id => %d\nname => %s\n", i.Id, i.Name)
}
