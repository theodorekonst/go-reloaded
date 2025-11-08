package transform

import (
	"strings"
	"unicode"

	"go-reloaded/internal/token"
)

func ApplyArticleAn(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		t := toks[i]
		if t.K == token.Word && (t.Text == "a" || t.Text == "A") {
			// find next Word (skip spaces and quotes)
			next := i + 1
			for next < len(toks) && (toks[next].K == token.Space || toks[next].K == token.Quote) {
				next++
			}
			if next < len(toks) && toks[next].K == token.Word {
				if startsWithVowelOrH(toks[next].Text) {
					if t.Text == "A" {
						t.Text = "An"
					} else {
						t.Text = "an"
					}
				}
			}
		}
		out = append(out, t)
	}
	return out
}

func startsWithVowelOrH(s string) bool {
	if s == "" {
		return false
	}
	r := []rune(s)[0]
	r = unicode.ToLower(r)
	return strings.ContainsRune("aeiouh", r)
}
