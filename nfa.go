package rscregexplearning

import "fmt"

type State struct {
	C        int
	Out      *State
	Out1     *State
	Lastlist int
}

func NewMatchState() *State {
	return &State{C: 257}
}

func NewSplitState(out, out1 *State) *State {
	return &State{C: 256, Out: out, Out1: out1}
}

type Frag struct {
	Start *State
	Out   *Ptrlist
}

type Ptrlist []**State

func list1(outp **State) Ptrlist {
	return []**State{outp}
}

func patch(l *Ptrlist, s *State) {
	for _, p := range *l {
		*p = s
	}
}

func Post2nfa(postfix string) *State {
	stack := make([]Frag, 0)
	push := func(s Frag) {
		stack = append(stack, s)
	}
	pop := func() Frag {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return s
	}
	for _, c := range postfix {
		print(c)
		switch {
		case c == '.':
			e2 := pop()
			e1 := pop()
			patch(e1.Out, e2.Start)
			push(Frag{Start: e1.Start, Out: e2.Out})
		case c == '|':
			e2 := pop()
			e1 := pop()
			s := NewSplitState(e1.Start, e2.Start)
			out := append(*e1.Out, *e2.Out...)
			push(Frag{Start: s, Out: &out})
		case c == '?':
			e := pop()
			s := NewSplitState(e.Start, nil)
			out := append(*e.Out, list1(&s.Out1)...)
			push(Frag{Start: s, Out: &out})
		case c == '*':
			e := pop()
			s := NewSplitState(e.Start, nil)
			patch(e.Out, s)
			out := list1(&s.Out1)
			push(Frag{Start: s, Out: &out})
		case c == '+':
			e := pop()
			s := NewSplitState(e.Start, nil)
			patch(e.Out, s)
			out := list1(&s.Out1)
			push(Frag{Start: e.Start, Out: &out})
		case c < 256:
			s := &State{C: int(c), Out: nil, Out1: nil}
			out := list1(&s.Out)
			push(Frag{Start: s, Out: &out})
		default:
			panic(fmt.Sprintf("invalid postfix: %d", c))
		}
	}

	e := pop()
	patch(e.Out, NewMatchState())
	return e.Start
}
