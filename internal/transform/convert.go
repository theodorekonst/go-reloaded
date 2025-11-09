package transform

import (
	"strconv"
	"strings"

	"go-reloaded/internal/token"
)

func isValidHex(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range strings.ToLower(s) {
		if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f')) {
			return false
		}
	}
	return true
}

func isValidBin(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r != '0' && r != '1' {
			return false
		}
	}
	return true
}

func ApplyHex(toks []token.Tok) []token.Tok {
	// First pass: modify tokens in place
	for i, t := range toks {
		if t.K == token.Tag && t.Text == "(hex)" {
			// Find previous word and convert
			for j := i - 1; j >= 0; j-- {
				if toks[j].K == token.Word {
					if isValidHex(toks[j].Text) {
						if val, err := strconv.ParseInt(toks[j].Text, 16, 64); err == nil {
							toks[j].Text = strconv.FormatInt(val, 10)
						}
					}
					break
				}
			}
		}
	}

	// Second pass: filter out hex tags
	out := []token.Tok{}
	for _, t := range toks {
		if t.K == token.Tag && t.Text == "(hex)" {
			continue
		}
		out = append(out, t)
	}

	return out
}

func ApplyBin(toks []token.Tok) []token.Tok {
	// First pass: modify tokens in place
	for i, t := range toks {
		if t.K == token.Tag && t.Text == "(bin)" {
			// Find previous word and convert
			for j := i - 1; j >= 0; j-- {
				if toks[j].K == token.Word {
					if isValidBin(toks[j].Text) {
						if val, err := strconv.ParseInt(toks[j].Text, 2, 64); err == nil {
							toks[j].Text = strconv.FormatInt(val, 10)
						}
					}
					break
				}
			}
		}
	}

	// Second pass: filter out bin tags
	out := []token.Tok{}
	for _, t := range toks {
		if t.K == token.Tag && t.Text == "(bin)" {
			continue
		}
		out = append(out, t)
	}

	return out
}
