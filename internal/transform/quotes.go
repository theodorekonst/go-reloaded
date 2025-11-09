package transform

import (
	"unicode"

	"go-reloaded/internal/token"
)

// ApplyQuotes tightens spaces *inside* each pair of single quotes.
//
// Strategy:
//   1) Find a quote pair.
//   2) Drop Space tokens immediately after the opening quote and immediately before the closing quote.
//   3) Join the interior tokens to a string, trim leading/trailing Unicode whitespace,
//      then re-tokenize that interior string and append it back.
//      This guarantees no stray boundary spaces even if earlier stages fused them into words.
//   4) Emit closing quote. Unmatched quotes are left unchanged.
func ApplyQuotes(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	i := 0

	for i < len(toks) {
		if toks[i].K != token.Quote {
			out = append(out, toks[i])
			i++
			continue
		}

		// Find closing quote linearly
		j := i + 1
		for j < len(toks) && toks[j].K != token.Quote {
			j++
		}
		if j >= len(toks) {
			// No closing quote -> keep as-is
			out = append(out, toks[i])
			i++
			continue
		}

		// Emit opening quote
		out = append(out, toks[i])

		// Strip Space tokens just inside the pair
		k := i + 1
		for k < j && toks[k].K == token.Space {
			k++
		}
		l := j - 1
		for l >= k && toks[l].K == token.Space {
			l--
		}

		// Re-emit the interior: join -> trim outer whitespace -> re-tokenize
		if k <= l {
			inner := token.Join(toks[k : l+1])
			inner = trimUnicodeOuter(inner)
			if inner != "" {
				innerToks := token.Tokenize(inner)
				out = append(out, innerToks...)
			}
		}

		// Emit closing quote and jump past it
		out = append(out, toks[j])
		i = j + 1
	}

	return out
}

// trimUnicodeOuter trims only leading/trailing Unicode whitespace runes.
func trimUnicodeOuter(s string) string {
	rs := []rune(s)
	start, end := 0, len(rs)-1
	for start <= end && unicode.IsSpace(rs[start]) {
		start++
	}
	for end >= start && unicode.IsSpace(rs[end]) {
		end--
	}
	if start > end {
		return ""
	}
	return string(rs[start : end+1])
}