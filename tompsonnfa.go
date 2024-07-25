package rscregexplearning

type List struct {
	S []*State
	N int
}

var listid = 0

func Match(start *State, s string) bool {
	var l1, l2 List
	clist := startlist(start, &l1)
	nlist := &l2
	for i := 0; i < len(s); i++ {
		step(clist, int(s[i]), nlist)
		clist, nlist = nlist, clist
	}
	return ismatch(clist)
}

func ismatch(l *List) bool {
	for i := 0; i < len(l.S); i++ {
		if l.S[i].C == 257 { // 257 is the match state
			return true
		}
	}
	return false
}

func addstate(l *List, s *State) {
	if s.Lastlist == listid {
		return
	}
	s.Lastlist = listid
	if s.C == 256 {
		addstate(l, s.Out)
		addstate(l, s.Out1)
		return
	}
	l.S = append(l.S, s)
	l.N++
}

func startlist(start *State, l *List) *List {
	listid++
	l.N = 0
	addstate(l, start)
	return l
}

func step(clist *List, c int, nlist *List) {
	listid++
	nlist.N = 0
	for i := 0; i < clist.N; i++ {
		s := clist.S[i]
		if s.C == c {
			addstate(nlist, s.Out)
		}
	}
}
