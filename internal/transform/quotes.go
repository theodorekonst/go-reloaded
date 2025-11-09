package transform

import "go-reloaded/internal/token"

// ApplyQuotes tightens spaces *inside* each pair of single quotes.
// - Removes any number of Space tokens immediately after the opening quote,
//   and immediately before the closing quote.
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

		// find next closing quote linearly
		j := i + 1
		for j < len(toks) && toks[j].K != token.Quote {
			j++
		}
		if j >= len(toks) {
			// no closing quote -> keep as-is
			out = append(out, toks[i])
			i++
			continue
		}

		// opening quote
		out = append(out, toks[i])

		// left-trim: drop ALL Space tokens immediately inside opening quote
		k := i + 1
		for k < j && toks[k].K == token.Space {
			k++
		}
		// right-trim: drop ALL Space tokens immediately inside closing quote
		l := j - 1
		for l >= k && toks[l].K == token.Space {
			l--
		}

		// emit interior if any (k..l inclusive)
		if k <= l {
			out = append(out, toks[k:l+1]...)
		}

		// closing quote
		out = append(out, toks[j])

		// continue after closing quote
		i = j + 1
	}
	return out
}