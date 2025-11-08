package transform

import (
	"strings"

	"go-reloaded/internal/token"
)

// Collapse consecutive plain spaces into a single " " while preserving any spaces
// that contain newlines exactly as-is (requirement: preserve line breaks).
func ApplySpaces(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	lastWasPlainSpace := false

	for _, t := range toks {
		if t.K != token.Space {
			out = append(out, t)
			lastWasPlainSpace = false
			continue
		}

		// t is Space
		if strings.ContainsRune(t.Text, '\n') {
			// Preserve newline spacing exactly; also reset the plain-space flag.
			out = append(out, t)
			lastWasPlainSpace = false
			continue
		}

		// Plain space(s) (no newline) → collapse to a single " "
		if lastWasPlainSpace {
			// skip extra plain spaces
			continue
		}
		out = append(out, token.Tok{K: token.Space, Text: " "})
		lastWasPlainSpace = true
	}

	return out
}
