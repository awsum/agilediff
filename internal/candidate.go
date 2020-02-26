package internal

import (
	"golang.org/x/net/html"
)

type Candidate struct {
	Score float64
	node  *html.Node
}
