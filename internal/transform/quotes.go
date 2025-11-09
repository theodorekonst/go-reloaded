package transform

import "go-reloaded/internal/token"

// ApplyQuotes tightens spaces *inside* each pair of single quotes.
// - Removes any number of Space tokens immediately after the opening quote,
//   and immediately before the closing quote.
// - Works when quotes are glued to punctuation/words (no spaces outside altered).
// - Handles empty quotes ("''") and multiple pairs in a row robustly.
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

		// Find the next quote (we don't try to infer nesting; we pair linearly)
		j := i + 1
		for j < len(toks) && toks[j].K != token.Quote {
			j++
		}
		if j >= len(toks) {
			// No closing quote: keep as-is
			out = append(out, toks[i])
			i++
			continue
		}

		// Emit opening quote
		out = append(out, toks[i])

		// Trim spaces just inside the quotes:
		// left-trim: skip all Space tokens after opening
		k := i + 1
		for k < j && toks[k].K == token.Space {
			k++
		}
		// right-trim: move l left over Space tokens before closing
		l := j - 1
		for l >= k && toks[l].K == token.Space {
			l--
		}

		// Emit interior range (if any content remains)
		if k <= l {
			out = append(out, toks[k:l+1]...)
		}

		// Emit closing quote
		out = append(out, toks[j])

		// Continue scan *after* the closing quote
		i = j + 1
	}

	return out
}