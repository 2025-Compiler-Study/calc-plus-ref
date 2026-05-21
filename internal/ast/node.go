package ast

type Node interface {
	String() string
	StringDepth(int) string
}
