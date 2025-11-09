package transform

import (
	"strings"

	"go-reloaded/internal/token"
)

func hasNewline(s string) bool { return strings.ContainsRune(s, '\n') }

func isAsciiPunctMark(t token.Tok) bool {
	if t.K != token.Punct && t.K != token.Group {
		return false
	}
	txt := t.Text
	// Groups like "..." or "!?" should already be token.Group with multi-chars;
	// treat them like a punct that should have one space after.
	if txt == "..." || txt == "!?" || txt == "?!" {
		return true
	}
	// Only these singles get "one space after" behavior:
	if len(txt) == 1 {
		switch txt[0] {
		case '.', ',', '!', '?', ':', ';':
			return true
		}
	}
	return false
}

func ApplyPunctuation(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		t := toks[i]

		if isAsciiPunctMark(t) {
			// Remove ALL plain spaces before punct
			for len(out) > 0 && out[len(out)-1].K == token.Space && !hasNewline(out[len(out)-1].Text) {
				out = out[:len(out)-1]
			}
			out = append(out, t)

			// After punctuation, ensure exactly one plain space,
			// unless newline follows or we're at EOF
			if i+1 < len(toks) {
				for (i+1) < len(toks) && toks[i+1].K == token.Space && !hasNewline(toks[i+1].Text) {
					i++
				}
				if (i+1) < len(toks) && toks[i+1].K == token.Space && hasNewline(toks[i+1].Text) {
					// newline-space follows: let it pass naturally
				} else {
					out = append(out, token.Tok{K: token.Space, Text: " "})
				}
			}
			continue
		}

		// For all other tokens (including em dashes —), just pass through.
		out = append(out, t)
	}
	return out
}
