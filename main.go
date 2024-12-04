package main

import "fmt"

func PrintAndReturn[T fmt.Stringer](t T) T {
	fmt.Println(t.String())
	return t
}

type printableInt int
type printableString string

func (i printableInt) String() string    { return fmt.Sprintf("%d", i) }
func (s printableString) String() string { return string(s) }

type A struct {
	a int
}
type B struct {
	b float64
}

func (a *A) String() string { return "A" }
func (a *B) String() string { return "A" }

func main() {
	_ = PrintAndReturn(printableInt(1))
	_ = PrintAndReturn(printableString("string"))

	_ = PrintAndReturn(&A{})
	_ = PrintAndReturn(&B{})
}
