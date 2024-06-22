// Exercise 6.5:
// The type of each word used by IntSet is uintbitCount, but bitCount-bit arithmetic may be
// inefficient on a 32-bit platform. Modify the program to use the uint type, which is
// the most efficient unsigned integer type for the platform. Instead of dividing by bitCount,
// define a constant holding the effective size of uint in bits, 32 or bitCount. You can use
// the perhaps too-clever expression 32 << (^uint(0) >> 63) for this purpose.

package intset

import (
	"bytes"
	"fmt"
)

const (
	bitCount = 32 << (^uint(0) >> 63)
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bitCount, uint(x%bitCount)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bitCount, uint(x%bitCount)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Exercise 6.2:
// Define a variadic (*IntSet).AddAll(...int) method that allows a list of values
// to be added, such as s.AddAll(1, 2, 3).
func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

// return the number of elements
func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		for j := 0; j < bitCount; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

// remove x from set
func (s *IntSet) Remove(x int) {
	word, bit := x/bitCount, uint(x%bitCount)
	if word <= len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// return copy of the set
func (s *IntSet) Copy() *IntSet {
	t := new(IntSet)
	t.words = make([]uint, len(s.words))
	copy(t.words, s.words)
	return t
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Exercise 6.3:
// (*IntSet).UnionWith computes the union of two sets using |,
// the word-parallel bitwise OR operator.
// Implement methods for IntersectWith, DifferenceWith, and SymmetricDifference for the
// corresponding set operations. (The symmetric difference of two sets contains the
// elements present in one set or the other but not both.)
func (s *IntSet) IntersectWith(t *IntSet) {
	i := 0
	for i < len(t.words) && i < len(s.words) {
		s.words[i] &= t.words[i]
		i++
	}
	for i < len(s.words) {
		s.words[i] = 0
		i++
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Exercise 6.4:
// Add a method Elems that returns a slice containing the elements of the set,
// suitable for iterating over with a range loop.
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < bitCount; j++ {
			if word&(1<<j) != 0 {
				elems = append(elems, i*bitCount+j)
			}

		}
	}
	return elems
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitCount; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitCount*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
