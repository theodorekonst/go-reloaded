package transform

import "go-reloaded/internal/token"

func ApplyQuotes(toks []token.Token) []token.Token {
	out := make([]token.Token, 0, len(toks))
	
	for i := 0; i < len(toks); i++ {
		if toks[i].K == token.Quote {
			// Find matching closing quote of same type
			quoteChar := toks[i].Text
			j := i + 1
			for j < len(toks) && !(toks[j].K == token.Quote && toks[j].Text == quoteChar) {
				j++
			}
			
			if j >= len(toks) {
				// No matching quote found - keep as is
				out = append(out, toks[i])
				continue
			}

			// Found matching pair - tighten content
			out = append(out, toks[i]) // opening quote

			// Trim spaces from content
			content := toks[i+1 : j]
			
			// Remove leading spaces
			start := 0
			for start < len(content) && content[start].K == token.Space {
				start++
			}
			
			// Remove trailing spaces
			end := len(content) - 1
			for end >= start && content[end].K == token.Space {
				end--
			}
			
			// Add trimmed content
			if start <= end {
				out = append(out, content[start:end+1]...)
			}
			
			out = append(out, toks[j]) // closing quote
			i = j // skip to after closing quote
		} else {
			out = append(out, toks[i])
		}
	}
	
	return out
}