package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/awsum/agilediff/internal"
)

func main() {
	id := flag.String("id", "make-everything-ok-button", "")

	flag.Parse()

	if flag.NArg() != 2 {
		log.Fatalf("usage: agilediff [-id=make-everything-ok-button] <original> <sample>\n")
	}

	original := reader(flag.Arg(0))
	sample := reader(flag.Arg(1))
	// no defers since everything is readonly, bound and fatal

	matcher, err := internal.NewMatcher(original, *id)
	if err != nil {
		log.Fatalf("failed to init matcher: %v", err)
	}

	candidates, err := matcher.Filter(sample)
	if err != nil {
		log.Fatalf("failed to filter sample: %v", err)
	}

	template, err := template.New("candidate").Parse(`
{{ .Path }}
{{- range $rule := .Passed }}
	V {{ $rule }}
{{- end }}
{{- range $rule := .Failed }}
	X {{ $rule }}
{{- end }}
`)
	if err != nil {
		panic(fmt.Sprintf("invalid template: %v", err))
	}

	for _, candidate := range candidates {
		err := template.Execute(os.Stdout, &candidate)
		if err != nil {
			log.Fatalf("failed to render template: %v", err)
		}
	}
}

func reader(fn string) io.ReadCloser {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", fn, err)
	}
	return f
}
