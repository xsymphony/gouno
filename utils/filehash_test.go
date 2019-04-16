package utils

import (
	"strings"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"123", "pmWkWSBCL51Bfkhn79xPuKBKHz__H6B-mY6G9_eieuM="},
		{"中文.docx", "Pl7sZt8A-li9d5he5IRJkq8rAMCdnX2g5v2B3_RdXgI="},
		{"", "47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU="},
	}
	for _, test := range tests {
		r := CalculateHash(strings.NewReader(test.input))
		if r != test.output {
			t.Errorf("CalculateHash('%s')->%s, But get: %s", test.input, test.output, r)
		}
	}
}
