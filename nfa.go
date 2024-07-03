package rscregexplearning

type State struct {
	C        int
	Out      *State
	Out1     *State
	Lastlist int
}

type Frag struct {
	Start *State
	Out   Ptrlist
}

type Ptrlist []*State

func list1(outp []*State) Ptrlist           { return nil }
func append(l1 Ptrlist, l2 Ptrlist) Ptrlist { return nil }

func Post2nfa(postfix string) *State { return nil }
