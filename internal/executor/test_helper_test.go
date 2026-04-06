package executor

import "calcPlus/internal/symbols"

func ConfigurePreset(e *Engine) {
	e.Variables = *symbols.NewScopedTable[int]()
	presets := map[string]int{
		"a": 10,
		"b": 20,
	}
	for k, v := range presets {
		_ = e.Variables.Register(k)
		_ = e.Variables.SetSymbol(k, v)
	}
}
