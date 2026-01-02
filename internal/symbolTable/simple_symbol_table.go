package symbolTable

import "fmt"

type SimpleSymbolTable struct {
	table map[string]int
}

func NewSimpleSymbolTable() *SimpleSymbolTable {
	return &SimpleSymbolTable{table: make(map[string]int)}
}

func (s *SimpleSymbolTable) Register(name string) error {
	if _, ok := s.table[name]; ok {
		return fmt.Errorf("variable %s already registered", name)
	}
	s.table[name] = 0
	return nil
}

func (s *SimpleSymbolTable) GetVariable(name string) (int, error) {
	val, ok := s.table[name]
	if !ok {
		return 0, fmt.Errorf("variable %s not found", name)
	}
	return val, nil
}

func (s *SimpleSymbolTable) SetVariable(name string, val int) error {
	if _, ok := s.table[name]; !ok {
		return fmt.Errorf("variable %s not found", name)
	}
	s.table[name] = val
	return nil
}
