package transform

import (
	"strings"
	"unicode"

	"go-reloaded/internal/token"
)

// ApplyCaseTags handles (up), (low), (cap) and (up, n)/(low, n)/(cap, n).
// It counts only previous Word tokens (skips punctuation/spaces/quotes).
func ApplyCaseTags(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))

	for i := 0; i < len(toks); i++ {
		t := toks[i]

		// Only handle Tag tokens here; pass everything else through.
		if t.K != token.Tag {
			out = append(out, t)
			continue
		}

		mode, n, ok := parseCaseTag(t.Text)
		if !ok {
			// Unknown tag -> keep it so other transforms get a chance.
			out = append(out, t)
			continue
		}

		// Apply to previous n Word tokens (skip non-words).
		j := len(out) - 1
		applied := 0
		for j >= 0 && applied < n {
			if out[j].K == token.Word {
				switch mode {
				case "up":
					out[j].Text = strings.ToUpper(out[j].Text)
				case "low":
					out[j].Text = strings.ToLower(out[j].Text)
				case "cap":
					out[j].Text = capWord(out[j].Text)
				}
				applied++
			}
			j--
		}
		// Drop the tag itself by not appending it.
	}
	return out
}

// parseCaseTag parses "(up)", "(low, 3)", "(cap,6)", etc.
func parseCaseTag(s string) (mode string, n int, ok bool) {
	s = strings.TrimSpace(s)
	if len(s) < 3 || s[0] != '(' || s[len(s)-1] != ')' {
		return "", 0, false
	}
	inner := strings.TrimSpace(s[1 : len(s)-1]) // content without parentheses
	if inner == "" {
		return "", 0, false
	}

	parts := strings.Split(inner, ",")
	mode = strings.ToLower(strings.TrimSpace(parts[0]))

	switch mode {
	case "up", "low", "cap":
		// valid
	default:
		return "", 0, false
	}

	n = 1
	if len(parts) > 1 {
		numStr := strings.TrimSpace(parts[1])
		val := 0
		for _, r := range numStr {
			if r < '0' || r > '9' {
				// invalid number -> treat as unknown tag
				return "", 0, false
			}
			val = val*10 + int(r-'0')
		}
		if val > 0 {
			n = val
		}
	}
	return mode, n, true
}

func capWord(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
