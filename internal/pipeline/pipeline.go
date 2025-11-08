package pipeline

import (
	"go-reloaded/internal/token"
	"go-reloaded/internal/transform"
)

func ProcessText(in string) string {
	toks := token.Tokenize(in)

	// Numbers first
	toks = transform.ApplyHex(toks)
	toks = transform.ApplyBin(toks)

	// Case tags
	toks = transform.ApplyCaseTags(toks)

	// Quotes tightening, then article
	toks = transform.ApplyQuotes(toks)
	toks = transform.ApplyArticleAn(toks)

	// NEW: collapse duplicate plain spaces (preserve line breaks)
	toks = transform.ApplySpaces(toks)

	// Finally, punctuation spacing rules
	toks = transform.ApplyPunctuation(toks)

	return token.Join(toks)
}