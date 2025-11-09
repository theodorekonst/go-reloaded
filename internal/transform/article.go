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
			// look ahead: skip spaces & quotes to next Word
			j := i + 1
			for j < len(toks) && (toks[j].K == token.Space || toks[j].K == token.Quote) {
				j++
			}
			if j < len(toks) && toks[j].K == token.Word {
				if startsWithVowelOrH(toks[j].Text) {
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