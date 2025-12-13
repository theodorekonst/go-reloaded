package transform

import (
	"strings"

	"go-reloaded/internal/token"
)

// ApplySpaces collapses consecutive plain spaces to one " ".
// It preserves any spaces containing newlines exactly.
// If trimEnds is true, removes leading/trailing plain spaces.
func ApplySpacesWithTrim(toks []token.Tok, trimEnds bool) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	lastWasSpace := false

	for _, t := range toks {
		if t.K != token.Space {
			out = append(out, t)
			lastWasSpace = false
			continue
		}
		// Handle newline spaces separately
		if strings.ContainsRune(t.Text, '\n') {
			out = append(out, t)
			lastWasSpace = false
			continue
		}
		// Skip consecutive plain spaces
		if lastWasSpace {
			continue
		}
		// Add single space
		out = append(out, token.Tok{K: token.Space, Text: " "})
		lastWasSpace = true
	}

	if !trimEnds || len(out) == 0 {
		return out
	}

	// Trim leading plain space
	if out[0].K == token.Space && !strings.ContainsRune(out[0].Text, '\n') {
		out = out[1:]
	}
	// Trim trailing plain space
	if len(out) > 0 {
		last := out[len(out)-1]
		if last.K == token.Space && !strings.ContainsRune(last.Text, '\n') {
			out = out[:len(out)-1]
		}
	}
	return out
}

// Backwards compatible wrapper (no trim)
func ApplySpaces(toks []token.Tok) []token.Tok { return ApplySpacesWithTrim(toks, false) }
