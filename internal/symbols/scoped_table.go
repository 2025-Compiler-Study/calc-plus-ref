package symbols

import "fmt"

type ScopedTable[T any] struct {
	scopes []SimpleTable[T]
}

func NewScopedTable[T any]() *ScopedTable[T] {
	return &ScopedTable[T]{
		scopes: []SimpleTable[T]{*NewSimpleTable[T]()},
	}
}

func (s *ScopedTable[T]) PushScope() {
	s.scopes = append(s.scopes, *NewSimpleTable[T]())
}

func (s *ScopedTable[T]) Register(name string) error {
	return s.scopes[len(s.scopes)-1].Register(name)
}

func (s *ScopedTable[T]) GetSymbol(name string) (T, error) {
	for i := len(s.scopes) - 1; i >= 0; i-- {
		val, err := s.scopes[i].GetSymbol(name)
		if err == nil {
			return val, nil
		}
	}
	return *new(T), fmt.Errorf("variable %s not found", name)
}

func (s *ScopedTable[T]) SetSymbol(name string, val T) error {
	for i := len(s.scopes) - 1; i >= 0; i-- {
		err := s.scopes[i].SetSymbol(name, val)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("variable %s not found", name)
}

func (s *ScopedTable[T]) PopScope() {
	s.scopes = s.scopes[:len(s.scopes)-1]
}
