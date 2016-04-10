// Copyright 2016 mikan.

package data

// IntSlice is type of int slice.
type IntSlice []int

// Add adds x to slice.
func (s *IntSlice) Add(x int) {
	*s = append(*s, x)
}

// ToSlice provides raw slice.
func (s *IntSlice) ToSlice() []int {
	return *s
}
