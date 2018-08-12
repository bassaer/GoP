package lib

import (
	"fmt"
)

type Item struct {
	Id   int
	Name string
}

func (i *Item) Info() string {
	return fmt.Sprintf("id => %d : name => %s", i.Id, i.Name)
}

func (i *Item) Show() {
	fmt.Println(i.Info())
}
