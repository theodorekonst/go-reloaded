package transform

import (
	"strings"

	"go-reloaded/internal/token"
)

// ApplyTightenQuoteEdges removes ONLY the immediate plain-space tokens
// that appear right after an opening quote or right before a closing quote.
// It preserves newlines and interior spacing.
//
// Examples it fixes safely:
//   "' hello'"   -> "'hello'"
//   "'hello '"   -> "'hello'"
//   "'  hello  '" stays "'hello  '" at the edges (interior spaces preserved)
func ApplyTightenQuoteEdges(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	i := 0
	for i < len(toks) {
		// If not a quote, just copy.
		if toks[i].K != token.Quote {
			out = append(out, toks[i])
			i++
			continue
		}

		// We have an opening quote in toks[i]. Always copy it.
		openIdx := len(out)
		out = append(out, toks[i])
		i++

		// Find the matching closing quote in the *input* sequence.
		j := i
		for j < len(toks) && toks[j].K != token.Quote {
			j++
		}
		if j >= len(toks) {
			// Unmatched: just stream through until end.
			for i < len(toks) {
				out = append(out, toks[i])
				i++
			}
			break
		}

		// We have a quoted region: toks[i..j-1] are inside, toks[j] is closing.

		// 1) Drop ALL immediate plain spaces after opening quote (no newline)
		for i < j && toks[i].K == token.Space && !strings.ContainsRune(toks[i].Text, '\n') {
			i++
		}

		// 2) Copy interior up to (but not including) any trailing plain spaces before closing
		k := j - 1
		for k >= i && toks[k].K == token.Space && !strings.ContainsRune(toks[k].Text, '\n') {
			k--
		}
		for p := i; p <= k; p++ {
			out = append(out, toks[p])
		}

		// 3) Emit the closing quote
		out = append(out, toks[j])

		// Continue after closing quote
		i = j + 1

		_ = openIdx // kept for clarity; not used further
	}
	return out
}