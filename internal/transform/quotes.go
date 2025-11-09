package transform

import "go-reloaded/internal/token"

func ApplyQuotes(toks []token.Token) []token.Token {
	result := make([]token.Token, 0, len(toks))
	i := 0
	
	for i < len(toks) {
		if toks[i].K == 2 { // Quote=2
			quoteChar := toks[i].Text
			start := i
			i++
			
			content := []token.Token{}
			for i < len(toks) && !(toks[i].K == 2 && toks[i].Text == quoteChar) {
				content = append(content, toks[i])
				i++
			}
			
			if i < len(toks) {
				result = append(result, toks[start])
				
				// Trim leading spaces (Space=1)
				for len(content) > 0 && content[0].K == 1 {
					content = content[1:]
				}
				
				// Remove spaces before punctuation and all trailing spaces
				clean := []token.Token{}
				for j := 0; j < len(content); j++ {
					// Skip spaces before punctuation
					if content[j].K == 1 && j+1 < len(content) && content[j+1].K == 3 {
						continue
					}
					clean = append(clean, content[j])
				}
				
				// Aggressively trim ALL trailing spaces
				for len(clean) > 0 && clean[len(clean)-1].K == 1 {
					clean = clean[:len(clean)-1]
				}
				
				result = append(result, clean...)
				result = append(result, toks[i])
				i++
			} else {
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
