package rscregexplearning

import (
	"reflect"
	"testing"
)

func TestPost2nfa(t *testing.T) {
	tests := []struct {
		postfix string
		want    *State
	}{
		{
			postfix: "a",
			want: &State{
				C: 'a',
			},
		},
		{
			postfix: "ab.",
			want: &State{
				C: 'a',
				Out: &State{
					C: 'b',
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.postfix, func(t *testing.T) {
			if got := Post2nfa(tt.postfix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post2nfa() = %v, want %v", got, tt.want)
			}
		})
	}
}
