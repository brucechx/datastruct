package undirectd_graph

type Stack struct {
	data []Item
}

func NewStack() *Stack{
	return &Stack{data:make([]Item, 0)}
}

func (s *Stack) Push(i Item){
	s.data = append(s.data, i)
}

func (s *Stack) Pop() Item{
	val := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]
	return val
}

func (s *Stack) IsEmpty() bool{
	return len(s.data) == 0
}
