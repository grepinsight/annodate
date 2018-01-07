package main

import (
	"testing"
)

func TestProcLine(t *testing.T) {
	got := procLine("2017-01-01,new year", ",", 1, false)
	want := "Sunday,2017-01-01,new year"
	if got != want {
		t.Errorf("expecting %q, got %q", want, got)
	}
}
