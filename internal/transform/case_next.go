package transform

import (
	"strings"

	"go-reloaded/internal/token"
)

func ApplyCaseNextMarker(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	var pending string

	for i := 0; i < len(toks); i++ {
		t := toks[i]
		if t.K == token.Tag && strings.HasPrefix(t.Text, caseNextMarkerPrefix) {
			pending = strings.TrimPrefix(t.Text, caseNextMarkerPrefix)
			continue // swallow marker
		}
		if pending != "" && t.K == token.Word {
			switch pending {
			case "up":
				t.Text = upWord(t.Text)
			case "low":
				t.Text = lowWord(t.Text)
			case "cap":
				t.Text = capWord(t.Text)
			}
			pending = "" // consume
		}
		out = append(out, t)
	}
	return out
}


