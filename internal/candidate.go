package internal

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Candidate gluing together processing pipeline and results presentation.
// Should not provide logic besides one that required for presentation.
type Candidate struct {
	Score  float64
	Failed []string
	Passed []string
	node   *html.Node
}

// Fail add test case as failed.
func (c *Candidate) Fail(rule string) {
	c.Failed = append(c.Failed, rule)
}

// Pass add test case as passed
func (c *Candidate) Pass(rule string) {
	c.Passed = append(c.Passed, rule)
}

// Path tries to build something simillar to xpath.
// NOTE: it's limited and broken, use only for presentation purposes
func (c *Candidate) Path() string {
	path := []string{}
	node := *c.node
	for ; node.Parent != nil; node = *node.Parent {
		path = append(path, fmt.Sprintf("%s%s", node.Data, domOrder(node)))
	}
	return strings.Join(path, " < ")
}

func domOrder(node html.Node) string {
	if node.Parent == nil {
		return ""
	}

	order := 0
	tag := node.Data
	for node.PrevSibling != nil {
		node = *node.PrevSibling
		if node.Data == tag {
			order += 1
		}
	}

	if order == 0 {
		return ""
	}

	return fmt.Sprintf("[%d]", order)
}
