package main

import (
	"fmt"
	"intset/intset"
)

func main() {
	var x, y intset.IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Len())

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	x.Remove(9)
	x.Remove(144)
	fmt.Println(x.String()) // "{1 42 }"

	x.Clear()
	fmt.Println(x.String()) // "{}"

	n := y.Copy()
	n.Add(577)
	y.Add(775)
	fmt.Println(n.String()) // {9 42 577}
	fmt.Println(y.String()) // {9 42 775}

	x.AddAll(1, 18, 43, 433, 9)
	fmt.Println(x.String())

	n.IntersectWith(&y)
	fmt.Println(n.String()) // {9 42}
	n.IntersectWith(&x)
	fmt.Println(n.String()) // {9}

	var s, t intset.IntSet

	s.AddAll(1, 18, 43, 433, 9)
	t.AddAll(1, 22, 45, 100, 9)
	s.IntersectWith(&t)
	fmt.Println(s.String())

	s.Clear()
	t.Clear()
	s.AddAll(1, 18, 43, 433, 9)
	t.AddAll(1, 22, 45, 100, 9)
	s.DifferenceWith(&t)
	fmt.Println(s.String())

	s.Clear()
	t.Clear()
	s.AddAll(1, 18, 43, 433, 9)
	t.AddAll(1, 22, 45, 100, 9)
	s.SymmetricDifference(&t)
	fmt.Println(s.String())

	fmt.Printf("{")
	for _, el := range s.Elems() {
		fmt.Printf(" %d", el)
	}
	fmt.Printf(" }\n")
}
