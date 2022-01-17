package main

import "testing"

func TestTruth(t *testing.T)  {
	if true != tru {
		t.Fatal("The world is crumbling")
	}
}