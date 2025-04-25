package lists

type Component struct {
	Value any
	next  *Component
}

func (c *Component) Next() *Component {
	return c.next
}

type Spisok struct {
	cmp Component
	len int
}

func (s *Spisok) Init() *Spisok {
	s.len = 0
	s.cmp.next = &s.cmp
	return s
}

func (s *Spisok) Front() *Component {
	return s.cmp.next
}

func (s *Spisok) PushBack(value any) *Component {
	s.cmp.next.Value = value
	return &s.cmp
}

func NewSpisok() *Spisok {
	return new(Spisok).Init()
}
