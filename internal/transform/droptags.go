package transform

import "go-reloaded/internal/token"

// ApplyDropTags removes any remaining Tag tokens (unknown/malformed).
func ApplyDropTags(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for _, t := range toks {
		if t.K == token.Tag {
			if t.Text == "()" {
				// keep literal ()
				out = append(out, token.Tok{K: token.Word, Text: "()"})
				continue
			}
			// otherwise drop it
			continue
		}
		out = append(out, t)
	}
	return out
}
