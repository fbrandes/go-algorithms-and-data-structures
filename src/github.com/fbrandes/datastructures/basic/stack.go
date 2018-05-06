package basic

type StackItem int

type Stack struct {
	size int
	items []StackItem
}

func (s *Stack) New() *Stack {
	s.size = 0
	s.items = []StackItem{}
	return s
}

func (s *Stack) Push(a StackItem) {
	s.items = append(s.items, a)
	s.size++
}

func (s *Stack) Pop() StackItem {
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.size--
	return item
}

func (s *Stack)Peek() StackItem {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpyt() bool {
	return len(s.items) == 0
}
