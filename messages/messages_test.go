package messages

import (
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"
	if got != expect {
		t.Errorf("did not get expected result. Wanted %q, got %q\n", expect, got)
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Goodbye, Gopher.\n"
	if got != expect {
		t.Errorf("did not get expected result. Wanted %q, got %q\n", expect, got)
	}
}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher!\n"},
		{input: "", expect: "Hello, !\n"},
	}
	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %q, got %q",
				s.input, got, s.expect)
		}
	}

}
