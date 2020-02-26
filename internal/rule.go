package internal

import "golang.org/x/net/html"

type Rule interface {
	Match(*html.Node) bool
}
