package main

import (
	"testing"
)

func TestProcLine(t *testing.T) {
	got := procLine("2018-01-01,new year", ",", 1, false)
	want := "Monday,2018-01-01,new year"
	if got != want {
		t.Errorf("expecting %q, got %q", want, got)
	}

	got2 := procLine("new year\t2018-01-01", "\t", 1, true)
	want2 := "new year\t2018-01-01\tMonday"
	if got2 != want2 {
		t.Errorf("expecting %q, got %q", want2, got2)
	}
}
