package internal

import (
	"fmt"
	"io"
	"sort"

	"github.com/PuerkitoBio/goquery"
	"github.com/awsum/agilediff/internal/rules"
	"golang.org/x/net/html"
)

// Matcher high level interface for node filtering, should not expose
// implementation details
type Matcher struct {
	rules map[string]Rule
}

func NewMatcher(input io.Reader, id string) (*Matcher, error) {
	// find golden element
	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input: %v", err)
	}
	selection := doc.Find(fmt.Sprintf("#%s", id))
	if len(selection.Nodes) != 1 {
		return nil, fmt.Errorf("malformed input: expected single golden element got %d", len(selection.Nodes))
	}
	node := selection.Nodes[0]

	// init default rules
	rules := map[string]Rule{
		"has onclick attribute": rules.NewHasOnclick(node),
		"same html tag":         rules.NewIsSameTag(node),
		"same dom depth":        rules.NewIsSameDepth(node),
		"same class":            rules.NewIsSameClass(node),
	}

	return &Matcher{
		rules: rules,
	}, nil
}

func (m *Matcher) Filter(sample io.Reader) ([]Candidate, error) {
	doc, err := goquery.NewDocumentFromReader(sample)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input: %v", err)
	}

	everyone := doc.Find("*")
	candidates := make([]Candidate, 0, len(everyone.Nodes))
	for _, node := range everyone.Nodes {
		candidates = append(candidates, m.report(node))
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Score > candidates[j].Score
	})
	// FIXME: len 0
	trimAt := 1
	for ; trimAt < len(candidates); trimAt++ {
		if candidates[trimAt].Score != candidates[trimAt-1].Score {
			break
		}
	}
	return candidates[:trimAt], nil
}

func (m *Matcher) report(node *html.Node) Candidate {
	candidate := Candidate{
		Score: 1,
		node:  node,
	}
	for name, rule := range m.rules {
		if !rule.Match(node) {
			candidate.Score *= 0.75
			candidate.Fail(name)
		} else {
			candidate.Pass(name)
		}
	}
	return candidate
}
