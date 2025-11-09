package transform

import "go-reloaded/internal/token"

// ValidateTags converts malformed tags (no space before) to punctuation
func ValidateTags(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		t := toks[i]
		if t.K == token.Tag {
			// Check if previous token is Space
			if i == 0 || toks[i-1].K != token.Space {
				// Malformed tag: convert to punctuation
				t.K = token.Punct
			}
		}
		out = append(out, t)
	}
	return out
}
