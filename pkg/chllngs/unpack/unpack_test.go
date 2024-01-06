package unpack

import (
	"errors"
	"testing"
)

var tests = []struct {
	packed, unpacked string
	err              error
}{
	{"a4bc2d5e", "aaaabccddddde", nil},
	{"abcd", "abcd", nil},
	{"a1", "a", nil},
	{"a0", "", NewErrZeroQuantifier('0', 1)},
	{"45", "", NewErrLonelyQuantifier('4', 0)},
	{"a2", "aa", nil},
	{"-", "", NewErrUnsupportedRune('-', 0)},
	{"a.", "", NewErrUnsupportedRune('.', 1)},
	{"\n.", "", NewErrUnsupportedRune('\n', 0)},
	{"qwe\\4\\5", "qwe45", nil},
	{"qwe\\45", "qwe44444", nil},
	{"qwe\\\\5", "qwe\\\\\\\\\\", nil},
	{"qwe\\a", "", NewErrInvalidEscapeSeq('a', 4)},
	{"\\", "", NewErrUnfinishedEscapeSeq('\\', 0)},
	{"", "", nil},
	{"\xf0\x28\x8c\x28", "", NewErrNotValidUtf8()},
	{"\xf8\xa1\xa1\xa1\xa1", "", NewErrNotValidUtf8()},
	{"\xe2\x82\x28", "", NewErrNotValidUtf8()},
	{"\xc3\x28", "", NewErrNotValidUtf8()},
	{"\xfc\xa1\xa1\xa1\xa1\xa1", "", NewErrNotValidUtf8()},
}

func TestSamples(t *testing.T) {
	for i, tcase := range tests {
		if result, err := Unpack(tcase.packed); !errors.Is(err, tcase.err) {
			t.Errorf("#%d test: errors don't match: %v !~ %v", i, err, tcase.err)
		} else if result != tcase.unpacked {
			t.Errorf("#%d test: result is invalid: %q != %q", i, result, tcase.unpacked)
		}
	}
}
