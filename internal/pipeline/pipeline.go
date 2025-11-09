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

	// Article correction
	toks = transform.ApplyArticleAn(toks)

	// Normalize spaces, then punctuation
	toks = transform.ApplySpaces(toks)
	toks = transform.ApplyPunctuation(toks)

	// Finally, quotes tightening (after punctuation spacing)
	toks = transform.ApplyQuotes(toks)

	// Final sweep: collapse any leftover plain spaces & trim ends
	toks = transform.ApplySpacesWithTrim(toks, true)

	return token.Join(toks)
}