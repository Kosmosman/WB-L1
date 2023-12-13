package set

type StringSet struct {
	Data map[string]struct{}
}

func CreateNewSet() *StringSet {
	return &StringSet{Data: make(map[string]struct{})}
}

func (s *StringSet) AddString(str string) {
	s.Data[str] = struct{}{}
}

func (s *StringSet) AddStrings(strings []string) {
	for _, str := range strings {
		s.AddString(str)
	}
}

func (s *StringSet) Erase(str string) {
	if s.Contains(str) {
		delete(s.Data, str)
	}
}

func (s *StringSet) Contains(str string) bool {
	_, ok := s.Data[str]
	return ok
}
