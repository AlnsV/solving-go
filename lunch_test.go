package solver

import (
	"container/list"
	"fmt"
	"testing"
)

func TestLunchBoxes(t *testing.T) {

	lst := list.New()
	boxes := []string{"0", "0", "1", "1", "1", "0", "1", "1"}
	for _, box := range boxes {
		lst.PushBack(box)
	}
	result := LunchBoxes(
		4,
		lst,
		[]string{"1", "1", "1", "0", "0", "0", "0", "0"},
	)
	fmt.Println(result)
}
