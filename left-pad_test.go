package leftpad

import "testing"

func TestLeftPad(t *testing.T) {
	if tl := LeftPad("foo", 5, ""); tl != "  foo" {
		t.Errorf("Should be ' foo', but is '%s'", tl)
	}
	if tl := LeftPad("foobar", 6, ""); tl != "foobar" {
		t.Errorf("Should be 'foobar', but is '%s'", tl)
	}
	if tl := LeftPad("1", 2, "0"); tl != "01" {
		t.Errorf("Should be '01', but is '%s'", tl)
	}
	if tl := LeftPad("1", 2, "-"); tl != "-1" {
		t.Errorf("Should be '-1', but is '%s'", tl)
	}
}
