package main

import "testing"

func TestInt(t *testing.T) {
	if (1 + 1) != 2 {
		t.Log("something goes wrong")
		t.Fail()
	}
}
