package utils

import (
	"testing"
	"strings"
)

func TestCalculateHash(t *testing.T) {
	var tests = []struct{
		input string
		output string
	} {
		{"123", "pmWkWSBCL51Bfkhn79xPuKBKHz//H6B+mY6G9/eieuM="},
		{"中文.docx", "Pl7sZt8A+li9d5he5IRJkq8rAMCdnX2g5v2B3/RdXgI="},
		{"", "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU="},
	}
	for _, test := range tests {
		r := CalculateHash(strings.NewReader(test.input))
		if r != test.output {
			t.Errorf("CalculateHash('%s')->%s, But get: %s", test.input, test.output, r)
		}
	}
}
