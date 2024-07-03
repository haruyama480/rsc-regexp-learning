package rscregexplearning

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPost2nfa(t *testing.T) {
	tests := []struct {
		postfix string
		want    *State
	}{
		{
			postfix: "a",
			want: &State{
				C:   'a',
				Out: NewMatchState(),
			},
		},
		{
			postfix: "ab.",
			want: &State{
				C: 'a',
				Out: &State{
					C:   'b',
					Out: NewMatchState(),
				},
			},
		},
		{
			postfix: "ab|",
			want: &State{
				C: 256,
				Out: &State{
					C:   'a',
					Out: NewMatchState(),
				},
				Out1: &State{
					C:   'b',
					Out: NewMatchState(),
				},
			},
		},
		{
			postfix: "a?",
			want: &State{
				C: 256,
				Out: &State{
					C:   'a',
					Out: NewMatchState(),
				},
				Out1: NewMatchState(),
			},
		},
		{
			postfix: "a*",
			want: func() *State {
				sp := &State{
					C: 256,
					Out: &State{
						C: 'a',
					},
					Out1: NewMatchState(),
				}
				sp.Out.Out = sp
				return sp
			}(),
		},
		{
			postfix: "a+",
			want: func() *State {
				a := &State{
					C: 'a',
					Out: &State{
						C:    256,
						Out1: NewMatchState(),
					},
				}
				a.Out.Out = a
				return a
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.postfix, func(t *testing.T) {
			got := Post2nfa(tt.postfix)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Post2nfa() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
