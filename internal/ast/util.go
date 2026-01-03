package ast

import "strings"

func indent(depth int) string {
	return strings.Repeat("    ", depth)
}
