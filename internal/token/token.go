package token

import (
	"unicode"
)

type Kind int

const (
	Word Kind = iota
	Space
	Quote
	Punct
	Group
	Tag
)

type Tok struct {
	K    Kind
	Text string
}

// Alias for compatibility
type Token = Tok
type TokenType = Kind

// Tokenize scans the input once and emits meaningful tokens.
// It avoids regex pitfalls and never mutates/"masks" content.
func Tokenize(s string) []Tok {
	r := []rune(s)
	var out []Tok
	i := 0
	n := len(r)

	emit := func(k Kind, start, end int) {
		out = append(out, Tok{K: k, Text: string(r[start:end])})
	}

	isWordRune := func(rr rune) bool {
		return unicode.IsLetter(rr) || unicode.IsDigit(rr)
	}

	isPunct := func(rr rune) bool {
		switch rr {
		case '.', ',', '!', '?', ':', ';', '—', '/':
			return true
		default:
			return false
		}
	}

	for i < n {
		ch := r[i]

		// 1) Quote (apostrophes and double quotes)
		if ch == '\'' {
			leftWord := i-1 >= 0 && isWordRune(r[i-1])
			rightWord := i+1 < n && isWordRune(r[i+1])
			if leftWord && rightWord {
				// embedded in word -> let word scanner handle it
				// fall through to word scanning
			} else {
				emit(Quote, i, i+1)
				i++
				continue
			}
		}
		if ch == '"' {
			emit(Quote, i, i+1)
			i++
			continue
		}

		// 2) Group ("...", "?!", "!?")
		if ch == '.' {
			// "..." group
			if i+2 < n && r[i+1] == '.' && r[i+2] == '.' {
				emit(Group, i, i+3)
				i += 3
				continue
			}
			// single '.' punct
			emit(Punct, i, i+1)
			i++
			continue
		}
		if ch == '!' && i+1 < n && r[i+1] == '?' {
			emit(Group, i, i+2)
			i += 2
			continue
		}
		if ch == '?' && i+1 < n && r[i+1] == '!' {
			emit(Group, i, i+2)
			i += 2
			continue
		}

		// 3) Single punctuation
		if isPunct(ch) {
			emit(Punct, i, i+1)
			i++
			continue
		}

		// 4) Spaces (one run, preserve newlines for later)
		if unicode.IsSpace(ch) {
			start := i
			i++
			for i < n && unicode.IsSpace(r[i]) {
				i++
			}
			emit(Space, start, i)
			continue
		}

		// 5) Tag: any balanced (...) — classify as a Tag token
		if ch == '(' {
			start := i
			j := i + 1
			for j < n && r[j] != ')' {
				j++
			}
			if j < n && r[j] == ')' {
				// We found a closing ')'
				emit(Tag, start, j+1)
				i = j + 1
				continue
			}
			// No closing ')': fall through to word scan
		}

		// 6) Word: consume word runes and embedded apostrophes
		if isWordRune(ch) {
			start := i
			for i < n {
				c := r[i]
				if isWordRune(c) {
					i++
					continue
				}
				// allow apostrophe inside word when both sides are word runes
				if c == '\'' && i+1 < n && i-1 >= start && isWordRune(r[i-1]) && isWordRune(r[i+1]) {
					i++ // consume apostrophe as part of word
					continue
				}
				// allow hyphen inside a word between word runes
				if c == '-' && i+1 < n && i-1 >= start && isWordRune(r[i-1]) && isWordRune(r[i+1]) {
					i++ // consume hyphen as part of word
					continue
				}
				break
			}
			emit(Word, start, i)
			continue
		}

		// 7) Fallback: consume one rune to avoid infinite loop
		i++
	}

	return out
}

// Join concatenates tokens back into a string exactly as stored in tokens.
// (Spacing/punctuation rules are handled by transforms, not here.)
func Join(toks []Tok) string {
	// small, allocation-friendly builder
	size := 0
	for _, t := range toks {
		size += len(t.Text)
	}
	b := make([]byte, 0, size)
	for _, t := range toks {
		b = append(b, t.Text...)
	}
	return string(b)
}
