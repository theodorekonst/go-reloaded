package transform

import "go-reloaded/internal/token"

func ApplyQuotes(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	i := 0
	for i < len(toks) {
		if toks[i].K != token.Quote {
			out = append(out, toks[i])
			i++
			continue
		}

		// Find the next quote (do not try to be clever about nesting; spec uses paired singles)
		j := i + 1
		for j < len(toks) && toks[j].K != token.Quote {
			j++
		}
		if j >= len(toks) {
			// No closing quote; keep as-is
			out = append(out, toks[i])
			i++
			continue
		}

		// Append opening quote
		out = append(out, toks[i])

		// Left-trim: skip spaces immediately after opening
		k := i + 1
		for k < j && toks[k].K == token.Space {
			k++
		}

		// Right-trim: move l left over spaces immediately before closing
		l := j - 1
		for l >= k && toks[l].K == token.Space {
			l--
		}

		// Emit interior (if any)
		if k <= l {
			out = append(out, toks[k:l+1]...)
		}

		// Append closing quote
		out = append(out, toks[j])

		// Jump past the closing quote
		i = j + 1
	}
	return out
}