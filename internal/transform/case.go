package transform

import (
	"strings"
	"unicode"

	"go-reloaded/internal/token"
)

// ApplyCaseTags updates words affected by (up), (low), (cap) and their (mode, n) forms.
// It applies to the LAST n previous Word tokens (skip Space/Quote/Punct/Group), not including the tag itself.
func ApplyCaseTags(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))

	for i := 0; i < len(toks); i++ {
		t := toks[i]
		if t.K != token.Tag {
			out = append(out, t)
			continue
		}

		mode, n, kind := parseCaseTagTri(t.Text)
		switch kind {
		case caseUnknown:
			// Not a case tag → keep for other transforms (hex/bin) or drop later
			out = append(out, t)
			continue
		case caseMalformed:
			// Looks like case tag but malformed → drop tag
			continue
		case caseOK:
			if n == 0 {
				// Explicit no-op → drop the tag
				continue
			}
			// Walk backward over OUT (tokens already emitted) and transform the last n Word tokens
			applied := 0
			j := len(out) - 1
			for j >= 0 && applied < n {
				switch out[j].K {
				case token.Word:
					switch mode {
					case "up":
						out[j].Text = strings.ToUpper(out[j].Text)
					case "low":
						out[j].Text = strings.ToLower(out[j].Text)
					case "cap":
						out[j].Text = capWord(out[j].Text)
					}
					applied++
				default:
					// skip Space/Quote/Punct/Group
				}
				j--
			}
			// Drop the tag itself
			continue
		}
	}
	return out
}

type caseKind int

const (
	caseUnknown caseKind = iota // not a case tag at all (e.g., hex/bin)
	caseMalformed               // looks like case tag but invalid format/number
	caseOK                      // valid case tag
)

// parseCaseTagTri parses "(up)", "(low, 3)", "(cap,0)" etc.
// Returns (mode, n, kind). If kind==caseUnknown, ignore; if caseMalformed, drop; if caseOK, apply.
func parseCaseTagTri(s string) (mode string, n int, kind caseKind) {
	s = strings.TrimSpace(s)
	if len(s) < 3 || s[0] != '(' || s[len(s)-1] != ')' {
		return "", 0, caseUnknown
	}
	inner := strings.TrimSpace(s[1 : len(s)-1])
	if inner == "" {
		return "", 0, caseMalformed
	}

	parts := strings.Split(inner, ",")
	mode = strings.ToLower(strings.TrimSpace(parts[0]))

	switch mode {
	case "up", "low", "cap":
		// looks like a case tag → now parse n (optional)
	default:
		return "", 0, caseUnknown
	}

	// default n = 1
	n = 1
	if len(parts) > 1 {
		numStr := strings.TrimSpace(parts[1])
		val := 0
		if numStr == "" {
			return "", 0, caseMalformed
		}
		for _, r := range numStr {
			if r < '0' || r > '9' {
				return "", 0, caseMalformed
			}
			val = val*10 + int(r-'0')
		}
		// allow n == 0 (no-op)
		n = val
	}
	return mode, n, caseOK
}

func capWord(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}