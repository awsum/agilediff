package rules

import "golang.org/x/net/html"

type IsSameClass struct {
	golden string
}

func NewIsSameClass(node *html.Node) *IsSameClass {
	return &IsSameClass{
		golden: attributeValue(node, "class"),
	}
}

func (r *IsSameClass) Match(node *html.Node) bool {
	return r.golden == attributeValue(node, "class")
}

func attributeValue(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
