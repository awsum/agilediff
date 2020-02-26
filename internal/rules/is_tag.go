package rules

import "golang.org/x/net/html"

type IsTag struct {
	golden string
}

func NewIsTag(node *html.Node) *IsTag {
	return &IsTag{
		golden: node.Data,
	}
}

func (r *IsTag) Match(node *html.Node) bool {
	return r.golden == node.Data
}
