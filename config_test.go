package main

import (
	"testing"
	"time"
)

func TestAtodu(t *testing.T) {
	var tests = []struct {
		input  string
		output time.Duration
	} {
		{"12", time.Duration(12)},
		{"13", time.Duration(13)},
		{"0", time.Duration(0)},
		{"A", time.Duration(0)},
		{"-1", time.Duration(-1)},
	}
	for _, test := range tests {
		o, _ := Atodu(test.input)
		if o != test.output {
			t.Errorf("Atodu(%v): %v; But get %v", test.input, test.output, o)
		}
	}
}
