package vm

type Op int

const (
	OpChar Op = iota
	OpMatch
	OpJmp
	OpSplit
)

type Inst struct {
	Op Op
	C  int
	X  int // index of prog
	Y  int // index of prog
}

type VM interface {
	Match(s string) bool
}
