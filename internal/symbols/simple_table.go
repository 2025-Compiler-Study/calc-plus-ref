package symbols

import "fmt"

type SimpleTable[T any] struct {
	table map[string]T
}

func NewSimpleTable[T any]() *SimpleTable[T] {
	return &SimpleTable[T]{table: make(map[string]T)}
}

func (s *SimpleTable[T]) Register(name string) error {
	if _, ok := s.table[name]; ok {
		return fmt.Errorf("variable %s already registered", name)
	}
	s.table[name] = *new(T)
	return nil
}

func (s *SimpleTable[T]) GetSymbol(name string) (T, error) {
	val, ok := s.table[name]
	if !ok {
		return *new(T), fmt.Errorf("variable %s not found", name)
	}
	return val, nil
}

func (s *SimpleTable[T]) SetSymbol(name string, val T) error {
	if _, ok := s.table[name]; !ok {
		return fmt.Errorf("variable %s not found", name)
	}
	s.table[name] = val
	return nil
}
