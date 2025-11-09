package transform

import (
	"strings"

	"go-reloaded/internal/token"
)

func CapitalizeI(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		t := toks[i]
		if t.K == token.Word {
			word := t.Text

			// Check if it's the standalone letter "i"
			if strings.ToLower(word) == "i" {
				t.Text = "I"
			} else if len(word) >= 2 && strings.ToLower(word[:1]) == "i" && word[1] == '\'' {
				// Check if it's "i'm", "i'll", "i've", etc.
				t.Text = "I" + word[1:]
			}
		}
		out = append(out, t)
	}
	return out
}
