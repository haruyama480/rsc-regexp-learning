package rscregexplearning

import (
	"testing"
)

func TestMatch(t *testing.T) {
	type args struct {
		regexp string
		s      string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"ab.", "ab"}, true},
		// {"", args{"ab.c.", "abc"}, true},
		{"", args{"ab|", "a"}, true},
		{"", args{"ab|", "b"}, true},
		// {"", args{"a+b+", "aabbb"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := Post2nfa(tt.args.regexp)
			// fmt.Printf("%v\n", start)
			if got := Match(start, tt.args.s); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
