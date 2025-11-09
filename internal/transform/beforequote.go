package transform

import "go-reloaded/internal/token"

// ApplySpaceBeforeOpeningQuote inserts a single plain space when a Word
// is immediately followed by an opening Quote with no space in between:
//   this'works' -> this 'works'
func ApplySpaceBeforeOpeningQuote(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		out = append(out, toks[i])

		// If current is Word and next is Quote, inject one space.
		if toks[i].K == token.Word && i+1 < len(toks) && toks[i+1].K == token.Quote {
			out = append(out, token.Tok{K: token.Space, Text: " "})
		}
	}
	return out
}