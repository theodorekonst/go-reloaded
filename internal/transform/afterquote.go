package transform

import "go-reloaded/internal/token"

// ApplySpaceAfterClosingQuote inserts one plain space *after* a closing quote
// when the next token is a Word (and there isn't already a space).
// It does nothing if the next token is punctuation or a newline Space.
func ApplySpaceAfterClosingQuote(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		out = append(out, toks[i])

		if toks[i].K == token.Quote {
			// If next is a Word (immediately), inject a single plain space.
			if i+1 < len(toks) && toks[i+1].K == token.Word {
				out = append(out, token.Tok{K: token.Space, Text: " "})
			}
		}
	}
	return out
}