package main

import "testing"

func TestBoo(t *testing.T) {
	if false {
		t.Error("Expected 1.5, got ")
	}
}
