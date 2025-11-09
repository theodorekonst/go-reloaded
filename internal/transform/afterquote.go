package transform

import "go-reloaded/internal/token"

// ApplySpaceAfterClosingQuote inserts a single plain space if a closing quote
// is immediately followed by a Word (with no space). It leaves punctuation/newlines alone.
func ApplySpaceAfterClosingQuote(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		out = append(out, toks[i])
		if toks[i].K == token.Quote {
			// If next is a Word (and not already spaced), insert one plain space.
			if i+1 < len(toks) && toks[i+1].K == token.Word {
				out = append(out, token.Tok{K: token.Space, Text: " "})
			}
		}
	}
	return out
}