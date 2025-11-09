package transform

import "go-reloaded/internal/token"

// ApplyDashQuoteTight removes a plain Space between an em dash (—) and a Quote.
// Example: "Then— 'done'" -> "Then—'done'"
func ApplyDashQuoteTight(toks []token.Tok) []token.Tok {
	out := make([]token.Tok, 0, len(toks))
	i := 0
	for i < len(toks) {
		// Need pattern: [em-dash] [Space] [Quote]
		if i+2 < len(toks) &&
			toks[i].K == token.Punct && toks[i].Text == "—" &&
			toks[i+1].K == token.Space &&
			toks[i+2].K == token.Quote {
			// emit dash, skip space, emit quote
			out = append(out, toks[i], toks[i+2])
			i += 3
			continue
		}
		out = append(out, toks[i])
		i++
	}
	return out
}