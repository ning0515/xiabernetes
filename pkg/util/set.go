package util

type empty struct {
}

type StringSet map[string]empty

func (s StringSet) Insert(items ...string) {
	for _, item := range items {
		s[item] = empty{}
	}
}

func (s StringSet) Delete(item string) {
	delete(s, item)
}
func (s StringSet) Has(item string) bool {
	_, contained := s[item]
	return contained
}

func NewStringSet(items ...string) StringSet {
	ss := StringSet{}
	ss.Insert(items...)
	return ss
}
