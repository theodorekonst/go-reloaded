package transform

import (
	"strings"

	"go-reloaded/internal/token"
)

func ApplyCaseNextMarker(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	var pendingMode string

	applyWord := func(s string) string {
		switch pendingMode {
		case "up":
			return strings.ToUpper(s)
		case "low":
			return strings.ToLower(s)
		case "cap":
			return capWord(s)
		default:
			return s
		}
	}

	for i := 0; i < len(toks); i++ {
		t := toks[i]
		if t.K == token.Tag && strings.HasPrefix(t.Text, caseNextMarkerPrefix) {
			// capture mode and skip emitting the marker
			pendingMode = strings.TrimPrefix(t.Text, caseNextMarkerPrefix)
			continue
		}

		if pendingMode != "" && t.K == token.Word {
			t.Text = applyWord(t.Text)
			pendingMode = "" // consume
		}

		out = append(out, t)
	}
	return out
}