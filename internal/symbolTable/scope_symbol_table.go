package symbolTable

import "fmt"

type ScopeSymbolTable struct {
	scopes []SimpleSymbolTable
}

func NewScopeSymbolTable() *ScopeSymbolTable {
	return &ScopeSymbolTable{
		scopes: []SimpleSymbolTable{*NewSimpleSymbolTable()},
	}
}

func (s *ScopeSymbolTable) PushScope() {
	s.scopes = append(s.scopes, *NewSimpleSymbolTable())
}

func (s *ScopeSymbolTable) Register(name string) error {
	return s.scopes[len(s.scopes)-1].Register(name)
}

func (s *ScopeSymbolTable) GetVariable(name string) (int, error) {
	for i := len(s.scopes) - 1; i >= 0; i-- {
		val, err := s.scopes[i].GetVariable(name)
		if err == nil {
			return val, nil
		}
	}
	return 0, fmt.Errorf("variable %s not found", name)
}

func (s *ScopeSymbolTable) SetVariable(name string, val int) error {
	for i := len(s.scopes) - 1; i >= 0; i-- {
		err := s.scopes[i].SetVariable(name, val)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("variable %s not found", name)
}

func (s *ScopeSymbolTable) PopScope() {
	s.scopes = s.scopes[:len(s.scopes)-1]
}
