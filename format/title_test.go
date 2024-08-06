package format_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-utils/format"
)

func Test_TitleCase(t *testing.T) {
	cases := [][]string{
		{"UNKNOWN", "Unknown"},
		{"WINDOWS_SERVER", "Windows Server"},
		{"UNDECIDED", "Undecided"},
		{"MALICIOUS", "Malicious"},
		{"BENIGN", "Benign"},
		{"NOT_APPLICABLE", "Not Applicable"},
	}
	for i := 0; i < len(cases); i++ {
		in := cases[i][0]
		exp := cases[i][1]
		t.Run(in, func(t *testing.T) {
			t.Parallel()
			result := format.TitleCase(in)
			assert.Equal(t, exp, result, in)
		})
	}
}
