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

	// Quotes: tighten, then ensure space after closing quote before words
	toks = transform.ApplyQuotes(toks)
	toks = transform.ApplySpaceAfterClosingQuote(toks)

	// Articles (a -> an)
	toks = transform.ApplyArticleAn(toks)

	// Normalize spaces, punctuation, drop any leftover tags, final tidy
	toks = transform.ApplySpaces(toks)
	toks = transform.ApplyPunctuation(toks)
	toks = transform.ApplyDropTags(toks)
	toks = transform.ApplySpacesWithTrim(toks, true)

	return token.Join(toks)
}