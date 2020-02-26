package rules

import "golang.org/x/net/html"

// IsSameTag matches nodes with same HTML tag (kinda) as golden node
type IsSameTag struct {
	golden string
}

func NewIsSameTag(node *html.Node) *IsSameTag {
	return &IsSameTag{
		golden: node.Data,
	}
}

func (r *IsSameTag) Match(node *html.Node) bool {
	return r.golden == node.Data
}
