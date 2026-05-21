package symbols

type Table[T any] interface {
	Register(name string) error
	GetSymbol(name string) (T, error)
	SetSymbol(name string, val T) error
}
