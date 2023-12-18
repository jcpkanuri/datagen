package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputs = []struct {
	wholeText string
	keyLen    int
	expected  string
	message   string
}{
	{"ABCDEFGH", 4, "ABCD", "First Four Chars"},
	{"ABCDEFGH", 8, "ABCDEFGH", "ALL"},
	{"", 8, "", "ALL"},
	{"A", 8, "A", "ALL"},
	{"2222222", 4, "2222", "First Four Chars"},
}

func TestTruncateText(t *testing.T) {

	s := "ABCDEFGHIJKLM"
	v1 := TruncateText(s, 1)

	if v1 != "A" {
		t.Fail()
	}

	for _, test := range testInputs {
		assert.Equal(t, test.expected, TruncateText(test.wholeText, test.keyLen))
	}

}
