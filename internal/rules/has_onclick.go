package rules

import "golang.org/x/net/html"

// HasOnclick matches nodes whose onclick attribute existense is the same as for golden one
type HasOnclick struct {
	golden bool
}

func NewHasOnclick(node *html.Node) *HasOnclick {
	return &HasOnclick{
		golden: hasAttribute(node, "onclick"),
	}
}

func (r *HasOnclick) Match(node *html.Node) bool {
	return r.golden == hasAttribute(node, "onclick")
}

func hasAttribute(node *html.Node, key string) bool {
	for _, attr := range node.Attr {
		if attr.Key == key && attr.Val != "" {
			return true
		}
	}
	return false
}
