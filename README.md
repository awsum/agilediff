# agilediff

HTML parsing will require maintenance no matter how smart heuristics are.
So focus for v1 was to provide something that is maintainable and open for extension with new rules and scoring mechanics.

## Usage

```bash
$ go run main.go -id=make-everything-ok-button testdata/sample-0-origin.html testdata/sample-1-evil-gemini.html
a[1] < div[1] < div < div < div[2] < div < div < body < html[1]
        V has onclick attribute
        V same html tag
        V same dom depth
        V same class
```