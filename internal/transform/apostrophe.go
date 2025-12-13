package transform

import "go-reloaded/internal/token"

// ApplyApostropheSpacing fixes all apostrophe spacing issues
func ApplyApostropheSpacing(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	
	for i := 0; i < len(toks); i++ {
		t := toks[i]
		
		if t.K == token.Quote && t.Text == "'" {
			// Possessive case: Word ' Word -> ensure no space before, space after
			if len(out) >= 1 && out[len(out)-1].K == token.Word {
				if i+1 < len(toks) && toks[i+1].K == token.Word {
					// This is possessive: James' car
					out = append(out, t)
					out = append(out, token.Tok{K: token.Space, Text: " "})
					continue
				}
			}
			// Remove space before possessive apostrophe: Word Space ' -> Word '
			if len(out) >= 2 && out[len(out)-1].K == token.Space && out[len(out)-2].K == token.Word {
				if i+1 < len(toks) && toks[i+1].K == token.Word {
					out = out[:len(out)-1] // remove space
					out = append(out, t)
					out = append(out, token.Tok{K: token.Space, Text: " "})
					continue
				}
			}
		}
		
		out = append(out, t)
	}
	
	return out
}