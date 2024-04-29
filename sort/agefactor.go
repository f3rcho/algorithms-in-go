package sort

import (
	"fmt"
	"sort"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
type Human struct {
	name string
	age  int
}

type AgeFactor []Human

func (a AgeFactor) Len() int {
	return len(a)
}

func (a AgeFactor) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a AgeFactor) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	audience := []Human{
		{"Alice", 35},
		{"Bob", 45},
		{"James", 25},
	}
	sort.Sort(AgeFactor(audience))
	fmt.Println(audience)
}
