package transform

import "go-reloaded/internal/token"

func ApplyQuotes(toks []token.Token) []token.Token {
	result := make([]token.Token, 0, len(toks))
	i := 0
	
	for i < len(toks) {
		if toks[i].K == token.Quote {
			// Find matching quote
			quoteChar := toks[i].Text
			start := i
			i++
			
			// Collect content between quotes
			content := []token.Token{}
			for i < len(toks) && !(toks[i].K == token.Quote && toks[i].Text == quoteChar) {
				content = append(content, toks[i])
				i++
			}
			
			if i < len(toks) { // Found closing quote
				// Add opening quote
				result = append(result, toks[start])
				
				// Add content without leading/trailing spaces
				for j, tok := range content {
					if tok.K == token.Space && (j == 0 || j == len(content)-1) {
						continue // Skip leading/trailing spaces
					}
					result = append(result, tok)
				}
				
				// Add closing quote
				result = append(result, toks[i])
				i++
			} else {
				// No closing quote found, add as-is
				result = append(result, toks[start])
				result = append(result, content...)
			}
		} else {
			result = append(result, toks[i])
			i++
		}
	}
	
	return result
}
