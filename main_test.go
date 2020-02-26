package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type Matcher struct {
	tag        string
	depth      int
	hasOnclick bool
}

func NewMatcher(node *html.Node) *Matcher {
	m := Matcher{
		tag:        node.Data,
		depth:      len(nodePath(node)),
		hasOnclick: hasOnclick(node),
	}
	return &m
}

func (m *Matcher) Score(node *html.Node) float64 {
	score := 1.
	if node.Data != m.tag {
		score *= 0.75
	}
	if m.depth != len(nodePath(node)) {
		score *= 0.75
	}
	if m.hasOnclick != hasOnclick(node) {
		score *= 0.75
	}
	return score
}

func hasOnclick(node *html.Node) bool {
	for _, attr := range node.Attr {
		if attr.Key == "onclick" {
			return true
		}
	}
	return false
}

func nodePath(node *html.Node) []string {
	path := []string{}
	for ; node.Parent != nil; node = node.Parent {
		// fmt.Println(node.Attr)
		path = append(path, fmt.Sprintf("%s", node.Data))
	}
	return path
}

// goquery to save a few lines over manually traversing html with net/html

func do(original, sample string) {
	input, err := os.Open(original)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		panic(err)
	}
	// doc.FindMatcher()
	selection := doc.Find("#make-everything-ok-button")
	node := selection.Nodes[0]
	matcher := NewMatcher(node)

	matcher = matcher
	{
		input, err := os.Open(sample)
		if err != nil {
			panic(err)
		}
		doc, err := goquery.NewDocumentFromReader(input)
		if err != nil {
			panic(err)
		}

		// iterate over all tags in document
		all := doc.Find("*")
		for _, node := range all.Nodes {
			score := matcher.Score(node)
			if score == 1 {
				fmt.Println(nodePath(node), matcher.Score(node), node.FirstChild.Data)
			}
		}
	}

}

func TestAll(t *testing.T) {
	cases := []struct {
		orig string
		samp string
	}{
		{
			"./testdata/sample-0-origin.html",
			"./testdata/sample-1-evil-gemini.html",
		},
		{
			"./testdata/sample-0-origin.html",
			"./testdata/sample-2-container-and-clone.html",
		},
		{
			"./testdata/sample-0-origin.html",
			"./testdata/sample-3-the-escape.html",
		},
		{
			"./testdata/sample-0-origin.html",
			"./testdata/sample-4-the-mash.html",
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s vs %s", tc.orig, tc.samp), func(t *testing.T) {
			do(tc.orig, tc.samp)
			t.Fail()
		})
	}
}
