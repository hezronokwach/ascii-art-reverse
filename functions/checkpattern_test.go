package reverse

import (
	"testing"
)

func TestCheckPattern(t *testing.T) {
	tests := []struct {
		char  []string
		word  []string
		want  bool
	}{
		{[]string{"pre", "sub"}, []string{"prefix", "submarine"}, true},
		{[]string{"pre", "sub"}, []string{"prefix", "subway"}, true},
		{[]string{"pre"}, []string{"prefix", "prelude"}, true},
		{[]string{"pre"}, []string{"suffix", "prefix"}, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := CheckPattern(tt.char, tt.word)
			if got != tt.want {
				t.Errorf("CheckPattern(%v, %v) = %v; want %v", tt.char, tt.word, got, tt.want)
			}
		})
	}
}

// TestTrimFound tests the TrimFound function.
func TestTrimFound(t *testing.T) {
	tests := []struct {
		length int
		word   []string
		want   []string
	}{
		
		{5, []string{"hello", "world"}, []string{"", "world"}},
		{0, []string{"test", "case"}, []string{"test", "case"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := TrimFound(tt.length, tt.word)
			if !equal(got, tt.want) {
				t.Errorf("TrimFound(%d, %v) = %v; want %v", tt.length, tt.word, got, tt.want)
			}
		})
	}
}

// Helper function to compare two slices of strings for equality.
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
