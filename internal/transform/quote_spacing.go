package transform

import (
	"go-reloaded/internal/token"
)

// ApplyQuoteSpacingFix removes spaces inside quotes and ensures proper spacing outside
func ApplyQuoteSpacingFix(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	
	for i := 0; i < len(toks); i++ {
		t := toks[i]
		
		// Handle quote tokens
		if t.K == token.Quote {
			// Find matching closing quote
			quoteChar := t.Text
			j := i + 1
			for j < len(toks) && (toks[j].K != token.Quote || toks[j].Text != quoteChar) {
				j++
			}
			
			if j < len(toks) {
				// Found matching quote pair
				out = append(out, t) // opening quote
				
				// Process interior tokens, removing spaces at edges
				interior := toks[i+1 : j]
				
				// Remove leading spaces
				start := 0
				for start < len(interior) && interior[start].K == token.Space {
					start++
				}
				
				// Remove trailing spaces
				end := len(interior) - 1
				for end >= start && interior[end].K == token.Space {
					end--
				}
				
				// Add cleaned interior
				if start <= end {
					out = append(out, interior[start:end+1]...)
				}
				
				// Add closing quote
				out = append(out, toks[j])
				
				// Ensure space after closing quote if followed by word
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