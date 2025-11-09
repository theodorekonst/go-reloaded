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

		// Check if this is "a" or "A" followed by a word starting with vowel sound
		if t.K == token.Word && (strings.ToLower(t.Text) == "a") {
			// Look ahead for the next word (skip spaces)
			j := i + 1
			for j < len(toks) && toks[j].K == token.Space {
				j++
			}

			if j < len(toks) && toks[j].K == token.Word {
				nextWord := strings.ToLower(toks[j].Text)
				if needsAn(nextWord) {
					// Change "a" to "an" preserving case
					if unicode.IsUpper(rune(t.Text[0])) {
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

func needsAn(word string) bool {
	if len(word) == 0 {
		return false
	}

	first := rune(word[0])
	
	// Simple vowel check
	switch first {
	case 'a', 'e', 'i', 'o':
		return true
	case 'u':
		// Special cases for 'u' sound
		if strings.HasPrefix(word, "uni") || strings.HasPrefix(word, "use") || strings.HasPrefix(word, "usual") {
			return false // "university", "user", "usual" use "a"
		}
		return true
	case 'h':
		// Special cases for 'h'
		if word == "hour" || word == "honest" || word == "honor" || word == "heir" {
			return true
		}
		return false
	default:
		return false
	}
}