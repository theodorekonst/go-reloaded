package transform

import "go-reloaded/internal/token"

// ApplyDropTags removes any remaining Tag tokens (unknown/malformed).
func ApplyDropTags(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for _, t := range toks {
		if t.K == token.Tag {
			// Drop it
			continue
		}
		out = append(out, t)
	}
	return out
}
