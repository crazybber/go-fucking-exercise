package skill

import (
	"fmt"
	"sort"
	"testing"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func sortByFunctions() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

type targetElements []struct {
	Key  int
	Name string
}

func (t targetElements) Len() int {
	return len(t)
}
func (t targetElements) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t targetElements) Less(i, j int) bool {
	return t[i].Key < t[j].Key
}

func TestSSortingByFunctions(t *testing.T) {
	sortByFunctions()

	targets := []struct {
		Key  int
		Name string
	}{
		{3, "crazy"},
		{2, "bber"},
		{-1, "negative"},
		{5, "bigger"},
	}

	pShow(targets)

	sort.Sort(targetElements(targets))

	assertEq(2, targets[1].Key)
	pShow(targets)
}
