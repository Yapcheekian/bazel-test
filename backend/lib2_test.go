package main

import "testing"

func TestAdd2(t *testing.T) {
	if add1(1, 2) != 3 {
		t.Fail()
	}
}
