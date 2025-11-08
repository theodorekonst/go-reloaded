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
			// --- BEFORE punctuation: remove ALL plain spaces (no newlines) ---
			for len(out) > 0 && out[len(out)-1].K == token.Space && !hasNewline(out[len(out)-1].Text) {
				out = out[:len(out)-1]
			}

			// Add the punctuation/group itself
			out = append(out, t)

			// --- AFTER punctuation: ensure exactly one plain space (unless newline follows) ---
			if i+1 < len(toks) {
				// 1) Skip ALL following spaces that are not newlines
				for (i+1) < len(toks) && toks[i+1].K == token.Space && !hasNewline(toks[i+1].Text) {
					i++
				}

				// 2) If the next token is a newline space, do NOT inject our own space
				if (i+1) < len(toks) && toks[i+1].K == token.Space && hasNewline(toks[i+1].Text) {
					// do nothing; the loop will append this newline space naturally
				} else {
					// Next is not a newline; insert exactly one plain space
					out = append(out, token.Tok{K: token.Space, Text: " "})
				}
			}
			continue
		}

		// default pass-through
		out = append(out, t)
	}
	return out
}
