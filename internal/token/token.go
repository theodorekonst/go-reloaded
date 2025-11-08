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

	isPunct := func(rr rune) bool {
		switch rr {
		case '.', ',', '!', '?', ':', ';':
			return true
		default:
			return false
		}
	}

	for i < n {
		ch := r[i]

		// 1) Quote
		if ch == '\'' {
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

		// 5) Tag: (letters [ , digits ] )
		if ch == '(' {
			start := i
			j := i + 1
			// tag name: [a-zA-Z]+
			for j < n && ((r[j] >= 'a' && r[j] <= 'z') || (r[j] >= 'A' && r[j] <= 'Z')) {
				j++
			}
			// optional part: , [spaces]* [digits]+
			if j < n && r[j] == ',' {
				j++
				for j < n && unicode.IsSpace(r[j]) {
					j++
				}
				dStart := j
				for j < n && r[j] >= '0' && r[j] <= '9' {
					j++
				}
				_ = dStart // we just allow; validation is done in transforms
			}
			// must end with ')'
			if j < n && r[j] == ')' {
				j++
				emit(Tag, start, j)
				i = j
				continue
			}
			// not a valid tag → fall back to word scan below
		}

		// 6) Word: consume until we hit space, quote, punct, group-start, tag-start, or parens
		start := i
		for i < n {
			c := r[i]
			// stop set: space, quote, punctuation, group starts, '(' or ')'
			if unicode.IsSpace(c) || c == '\'' || isPunct(c) || c == '(' || c == ')' ||
				c == '.' || c == '!' || c == '?' {
				break
			}
			i++
		}
		if start == i {
			// safety net: if we didn't advance (rare), consume one rune to avoid infinite loop
			i++
		} else {
			emit(Word, start, i)
		}
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