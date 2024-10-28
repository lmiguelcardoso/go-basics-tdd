package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("a")
	expected := "aaaaa"

	if repeat != expected{
		t.Errorf("expected %v got %v",expected, repeat)
	}
}