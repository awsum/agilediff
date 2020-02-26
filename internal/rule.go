package internal

import "golang.org/x/net/html"

// Rule basic heuristic
type Rule interface {
	// Match should succeed only if provided node matches golden according
	// to heuristic implementation
	Match(node *html.Node) bool
}
