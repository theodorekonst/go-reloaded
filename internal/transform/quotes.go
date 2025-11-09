package transform

import (
	"unicode"

	"go-reloaded/internal/token"
)

// ApplyQuotes tightens spaces *inside* each pair of single quotes.
// - Removes any number of Space tokens immediately inside the opening/closing quote.
// - Also trims leading/trailing Unicode whitespace on the first/last interior tokens.
// - Works when quotes are glued to punctuation/words (outside tokens unaffected).
// - Handles empty quotes and multiple pairs in a row robustly.
// - Leaves unmatched single quotes untouched.
func ApplyQuotes(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	i := 0
	for i < len(toks) {
		if toks[i].K != token.Quote {
			out = append(out, toks[i])
			i++
			continue
		}

		// Find the next closing quote linearly.
		j := i + 1
		for j < len(toks) && toks[j].K != token.Quote {
			j++
		}
		if j >= len(toks) {
			// No closing quote: keep as-is.
			out = append(out, toks[i])
			i++
			continue
		}

		// Emit opening quote.
		out = append(out, toks[i])

		// Left-trim: drop ALL Space tokens immediately inside opening quote.
		k := i + 1
		for k < j && toks[k].K == token.Space {
			k++
		}
		// Right-trim: drop ALL Space tokens immediately inside closing quote.
		l := j - 1
		for l >= k && toks[l].K == token.Space {
			l--
		}

		// Emit interior if any (k..l inclusive), with extra safety:
		if k <= l {
			// Trim leading whitespace runes on the first interior token text.
			if toks[k].K == token.Word || toks[k].K == token.Punct || toks[k].K == token.Group {
				toks[k].Text = trimLeftSpaces(toks[k].Text)
			}
			// Trim trailing whitespace runes on the last interior token text.
			if toks[l].K == token.Word || toks[l].K == token.Punct || toks[l].K == token.Group {
				toks[l].Text = trimRightSpaces(toks[l].Text)
			}
			// If either end became empty (paranoia), skip it.
			if k <= l {
				for x := k; x <= l; x++ {
					if toks[x].Text != "" {
						out = append(out, toks[x])
					}
				}
			}
		}

		// Emit closing quote.
		out = append(out, toks[j])

		// Continue after closing quote.
		i = j + 1
	}
	return out
}

func trimLeftSpaces(s string) string {
	// Trim any unicode whitespace at the front, if present.
	r := []rune(s)
	idx := 0
	for idx < len(r) && unicode.IsSpace(r[idx]) {
		idx++
	}
	return string(r[idx:])
}

func trimRightSpaces(s string) string {
	// Trim any unicode whitespace at the end, if present.
	r := []rune(s)
	idx := len(r) - 1
	for idx >= 0 && unicode.IsSpace(r[idx]) {
		idx--
	}
	if idx < 0 {
		return ""
	}
	return string(r[:idx+1])
}