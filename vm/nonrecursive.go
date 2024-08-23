package vm

type NonRecursiveBacktrack struct {
	prog []Inst
}

var _ VM = (*NonRecursiveBacktrack)(nil)

func NewNonRecursiveBacktrack(prog []Inst) *NonRecursiveBacktrack {
	return &NonRecursiveBacktrack{prog: prog}
}

type Thread struct {
	PC int
	S  string
}

func (vm *NonRecursiveBacktrack) Match(s string) bool {
	return false
}
