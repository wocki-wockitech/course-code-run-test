package main

import "testing"

func TestConcat(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{[]string{"a", "b", "c"}, "abc"},
		{[]string{}, ""},
		{[]string{"hello"}, "hello"},
		{[]string{"foo", " ", "bar"}, "foo bar"},
	}
	for _, tt := range tests {
		got := Concat(tt.input)
		if got != tt.want {
			t.Errorf("Concat(%v) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func BenchmarkConcat(b *testing.B) {
	parts := []string{"the", " ", "quick", " ", "brown", " ", "fox"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Concat(parts)
	}
}
