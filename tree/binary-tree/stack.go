package binary_tree

type Item interface {}

type Stack struct {
	Items []Item
}

func (s *Stack)New() *Stack{
	s.Items = []Item{}
	return s
}

// 压栈
func (s *Stack)Push(i Item)  {
	s.Items = append(s.Items, i)
}

// 出栈
func (s *Stack)Pop() Item {
	if s.isEmpty(){
		panic("栈为空")
	}
	item := s.Items[len(s.Items) - 1]
	s.Items = s.Items[:len(s.Items) - 1]
	return item
}

// isEmpty
func (s *Stack)isEmpty() bool {
	return len(s.Items) == 0
}
