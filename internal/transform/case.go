package transform

import (
	"strings"
	"unicode"

	"go-reloaded/internal/token"
)

const caseNextMarkerPrefix = "CASE_NEXT:"

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
			// Looks like case tag but malformed → keep original tag
			out = append(out, t)
			continue
		case caseOK:
			if n == 0 {
				// Explicit no-op → drop the tag
				continue
			}
			// 1) collect LAST n Word tokens from OUT using improved helper
			idxs := collectPreviousWordIdxsSameLine(out, len(out), n)

			// 3) Apply the transform
			applyWord := func(s string) string {
				switch mode {
				case "up":
					return upWord(s)
				case "low":
					return lowWord(s)
				case "cap":
					return capWord(s)
				default:
					return s
				}
			}

			// Apply to collected previous words only (no spill for compatibility)
			for _, k := range idxs {
				// Skip headers that start with "section"
				if strings.HasPrefix(strings.ToLower(strings.TrimSpace(out[k].Text)), "section") {
					continue
				}
				out[k].Text = applyWord(out[k].Text)
			}

			continue
		}
	}
	return out
}

type caseKind int

const (
	caseUnknown   caseKind = iota // not a case tag at all (e.g., hex/bin)
	caseMalformed                 // looks like case tag but invalid format/number
	caseOK                        // valid case tag
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
		return "", 0, caseUnknown // empty tags should be ignored, not malformed
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
	parts := strings.Split(s, "-")
	for i, p := range parts {
		if len(p) == 0 {
			continue
		}
		runes := []rune(p)
		runes[0] = unicode.ToTitle(runes[0]) // Title-case first rune
		for j := 1; j < len(runes); j++ {
			runes[j] = unicode.ToLower(runes[j])
		}
		parts[i] = string(runes)
	}
	return strings.Join(parts, "-")
}

func upWord(s string) string {
	// Uppercase all segments, keep hyphens
	parts := strings.Split(s, "-")
	for i, p := range parts {
		parts[i] = strings.ToUpper(p)
	}
	return strings.Join(parts, "-")
}

func lowWord(s string) string {
	parts := strings.Split(s, "-")
	for i, p := range parts {
		parts[i] = strings.ToLower(p)
	}
	return strings.Join(parts, "-")
}

func collectPreviousWordIdxsSameLine(toks []token.Tok, i, n int) []int {
	idxs := []int{}
	seen := 0
	for j := i - 1; j >= 0 && seen < n; j-- {
		if toks[j].K == token.Word {
			idxs = append(idxs, j)
			seen++
		}
		// Stop only at NEWLINES (not punctuation or spaces)
		if toks[j].K == token.Group && toks[j].Text == "\n" {
			break
		}
	}
	// reverse so they stay left→right
	for l, r := 0, len(idxs)-1; l < r; l, r = l+1, r-1 {
		idxs[l], idxs[r] = idxs[r], idxs[l]
	}
	return idxs
}
