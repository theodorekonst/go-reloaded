package transform

import (
	"strings"

	"go-reloaded/internal/token"
)

func hasNewline(s string) bool {
	return strings.ContainsRune(s, '\n')
}

func ApplyPunctuation(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		t := toks[i]

		if t.K == token.Punct || t.K == token.Group {
			// Remove ALL plain spaces before punct
			for len(out) > 0 && out[len(out)-1].K == token.Space && !hasNewline(out[len(out)-1].Text) {
				out = out[:len(out)-1]
			}
			out = append(out, t)

			// After punctuation, ensure exactly one plain space
			// unless newline follows, OR unless this is end-of-file.
			if i+1 < len(toks) {
				// Skip all plain spaces
				for (i+1) < len(toks) && toks[i+1].K == token.Space && !hasNewline(toks[i+1].Text) {
					i++
				}
				// If next is newline space, let it pass naturally; else inject one space
				if (i+1) < len(toks) && toks[i+1].K == token.Space && hasNewline(toks[i+1].Text) {
					// no-op
				} else {
					out = append(out, token.Tok{K: token.Space, Text: " "})
				}
			}
			continue
		}

		out = append(out, t)
	}
	return out
}
