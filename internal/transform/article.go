package transform

import (
	"strings"

	"go-reloaded/internal/token"
)

func ApplyArticleAn(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	for i := 0; i < len(toks); i++ {
		t := toks[i]
		if t.K == token.Word {
			article := strings.ToLower(t.Text)
			if article == "a" || article == "an" {
				// Find next WORD token (skip spaces, punctuation, quotes)
				nextWordIdx := -1
				for j := i + 1; j < len(toks); j++ {
					if toks[j].K == token.Word {
						nextWordIdx = j
						break
					}
				}
				
				if nextWordIdx != -1 {
					nextWord := toks[nextWordIdx].Text
					shouldBeAn := needsAn(nextWord)
					
					// Fix the article
					if shouldBeAn && article == "a" {
						t.Text = preserveCase("an", t.Text)
					} else if !shouldBeAn && article == "an" {
						t.Text = preserveCase("a", t.Text)
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
	
	lower := strings.ToLower(word)
	first := lower[0]
	
	// Special cases: words starting with vowel letters but consonant sounds
	if first == 'u' {
		// "university", "european", "unicorn" etc. start with 'y' sound
		if strings.HasPrefix(lower, "uni") || strings.HasPrefix(lower, "eu") {
			return false // keep "a"
		}
		return true // "umbrella", "uncle" need "an"
	}
	
	// Standard vowels
	if first == 'a' || first == 'e' || first == 'i' || first == 'o' {
		return true
	}
	
	// Silent 'h' words need "an"
	if first == 'h' {
		silentH := []string{"hour", "honest", "honor", "heir"}
		for _, word := range silentH {
			if strings.HasPrefix(lower, word) {
				return true
			}
		}
	}
	
	return false
}

func preserveCase(newWord, oldWord string) string {
	if len(oldWord) > 0 && oldWord[0] >= 'A' && oldWord[0] <= 'Z' {
		return strings.Title(newWord)
	}
	return newWord
}
