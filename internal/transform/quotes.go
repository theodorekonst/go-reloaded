package transform

import "go-reloaded/internal/token"

func ApplyQuotes(toks []token.Token) []token.Token {
	out := make([]token.Token, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		if toks[i].K == token.Quote {
			// Find matching closing quote
			j := i + 1
			for j < len(toks) && toks[j].K != token.Quote {
				j++
			}
			if j >= len(toks) {
				// No closing quote → keep as-is
				out = append(out, toks[i])
				continue
			}

			// Tighten: remove spaces right inside quotes
			// opening '
			out = append(out, toks[i])

			// left trim (right after opening)
			k := i + 1
			for k < j && toks[k].K == token.Space {
				k++
			}
			// right trim (right before closing)
			l := j - 1
			for l > k && toks[l].K == token.Space {
				l--
			}
			if k <= l {
				out = append(out, toks[k:l+1]...)
			}
			// closing '
			out = append(out, toks[j])

			i = j
			continue
		}
		out = append(out, toks[i])
	}
	return out
}
