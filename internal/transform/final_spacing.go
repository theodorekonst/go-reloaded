package transform

import "go-reloaded/internal/token"

// ApplyFinalSpacingFix handles all remaining spacing edge cases
func ApplyFinalSpacingFix(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	
	for i := 0; i < len(toks); i++ {
		t := toks[i]
		
		// Handle quote pairs to remove internal spaces
		if t.K == token.Quote {
			// Find matching closing quote
			quoteChar := t.Text
			j := i + 1
			for j < len(toks) && (toks[j].K != token.Quote || toks[j].Text != quoteChar) {
				j++
			}
			
			if j < len(toks) {
				// Found matching quote pair - ensure proper spacing
				// Add space before opening quote if needed
				if len(out) > 0 && (out[len(out)-1].K == token.Word || out[len(out)-1].K == token.Punct) {
					out = append(out, token.Tok{K: token.Space, Text: " "})
				}
				
				out = append(out, t) // opening quote
				
				// Process interior, removing edge spaces
				interior := toks[i+1 : j]
				start := 0
				for start < len(interior) && interior[start].K == token.Space {
					start++
				}
				end := len(interior) - 1
				for end >= start && interior[end].K == token.Space {
					end--
				}
				if start <= end {
					out = append(out, interior[start:end+1]...)
				}
				
				out = append(out, toks[j]) // closing quote
				
				// Add space after closing quote if followed by word
				if j+1 < len(toks) && toks[j+1].K == token.Word {
					out = append(out, token.Tok{K: token.Space, Text: " "})
				}
				
				i = j // skip to after closing quote
				continue
			}
		}
		
		out = append(out, t)
	}
	
	return out
}