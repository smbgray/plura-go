package main_test

import "testing"

func TestAdding(t *testing.T) {
	got := 3
	expected := 2
	if got != expected {
		t.Error("Test didnt pass")
	}
}
