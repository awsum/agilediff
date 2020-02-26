package rules

import (
	"golang.org/x/net/html"
)

// IsSameDepth matches nodes with same DOM depth as golden node
type IsSameDepth struct {
	golden int
}

func NewIsSameDepth(node *html.Node) *IsSameDepth {
	return &IsSameDepth{
		golden: nodeDepth(node),
	}
}

func (r *IsSameDepth) Match(node *html.Node) bool {
	return r.golden == nodeDepth(node)
}

func nodeDepth(node *html.Node) int {
	depth := 0
	for ; node.Parent != nil; node = node.Parent {
		depth += 1
	}
	return depth
}
