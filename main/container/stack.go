package container

type Stack struct {
	stack []interface{}
}

func (s *Stack) Push(i interface{}) {
	s.stack = append(s.stack, i)
}

func (s *Stack) IsEmpty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() interface{} {
	var c interface{}
	if s.IsEmpty() {
		return c
	}
	c = s.stack[len(s.stack) - 1]
	s.stack = s.stack[:len(s.stack) - 1]
	return c
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func NewStack() *Stack  {
	return &Stack{stack:make([]interface{}, 0)}
}