package vm

type RecursiveBacktrack struct {
	prog []Inst
}

var _ VM = (*RecursiveBacktrack)(nil)

func NewRecursiveBacktrack(prog []Inst) *RecursiveBacktrack {
	return &RecursiveBacktrack{prog: prog}
}

func (vm *RecursiveBacktrack) Match(s string) bool {
	r1 := vm.recursive(0, s)
	r2 := vm.recursiveloop(0, s)
	if r1 != r2 {
		panic("not equal")
	}
	return r1
}

func (vm *RecursiveBacktrack) recursive(pc int, s string) bool {
	cur := vm.prog[pc]
	switch cur.Op {
	case OpChar:
		if int(s[0]) != cur.C {
			return false
		}
		return vm.recursive(pc+1, s[1:])
	case OpMatch:
		return true
	case OpJmp:
		return vm.recursive(cur.X, s)
	case OpSplit:
		return vm.recursive(cur.X, s) || vm.recursive(cur.Y, s)
	}
	panic("unreachable")
}

func (vm *RecursiveBacktrack) recursiveloop(pc int, s string) bool {
	cur := vm.prog[pc]
	for {
		switch cur.Op {
		case OpChar:
			if int(s[0]) != cur.C {
				return false
			}
			cur = vm.prog[pc+1]
			s = s[1:]
			continue
		case OpMatch:
			return true
		case OpJmp:
			cur = vm.prog[cur.X]
			continue
		case OpSplit:
			if vm.recursiveloop(cur.X, s) {
				return true
			}
			cur = vm.prog[cur.Y]
			continue
		}
		panic("unreachable")
	}
}
