package sorting

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
type Human struct {
	Name string
	Age  int
}

type AgeFactor []Human

func (a AgeFactor) Len() int {
	return len(a)
}

func (a AgeFactor) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a AgeFactor) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
