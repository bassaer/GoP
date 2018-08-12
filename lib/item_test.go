package lib

import "testing"

func TestInfo(t *testing.T) {
	item := Item{Id: 100, Name: "test"}
	actual := item.Info()
	expected := "id => 100 : name => test"
	if actual != expected {
		t.Errorf("actual : %s\n expected : %s", actual, expected)
	}
}
