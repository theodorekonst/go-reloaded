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
					if shouldBeAn && article == "a" {
						if t.Text == "a" {
							t.Text = "an"
						} else {
							t.Text = "An"
						}
					} else if !shouldBeAn && article == "an" {
						if t.Text == "an" {
							t.Text = "a"
						} else {
							t.Text = "A"
						}
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

	// Standard vowels (except special cases)
	if first == 'a' || first == 'i' {
		return true
	}

	if first == 'e' {
		// "European" starts with 'y' sound
		if strings.HasPrefix(lower, "eu") {
			return false // keep "a"
		}
		return true
	}

	// Silent 'h' words need "an"
	if first == 'h' {
		silentH := []string{"hour", "honest", "honor", "heir", "herb"}
		for _, word := range silentH {
			if strings.HasPrefix(lower, word) {
				return true
			}
		}
	}

	// Numbers starting with vowel sounds but written as consonants
	if first == 'o' {
		// "one" starts with 'w' sound
		if strings.HasPrefix(lower, "one") {
			return false // keep "a"
		}
		return true
	}

	return false
}

func preserveCase(newWord, oldWord string) string {
	if len(oldWord) == 0 {
		return newWord
	}

	// Check if the old word is all uppercase AND longer than 1 char
	if len(oldWord) > 1 && strings.ToUpper(oldWord) == oldWord {
		return strings.ToUpper(newWord)
	}

	// Check if the old word starts with uppercase (including single "A")
	if oldWord[0] >= 'A' && oldWord[0] <= 'Z' {
		return strings.Title(newWord)
	}

	return newWord
}
