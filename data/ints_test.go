// Copyright 2016 mikan.

package data

import "testing"

func TestAdd(t *testing.T) {
	var s IntSlice
	s.Add(1)
	s.Add(2)
	if s.ToSlice()[0] != 1 {
		t.Errorf("expected %v but got %v", 1, s.ToSlice()[1])
	}
	if s.ToSlice()[1] != 2 {
		t.Errorf("expected %v but got %v", 2, s.ToSlice()[2])
	}
}
