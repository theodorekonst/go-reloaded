package pipeline

import (
	"go-reloaded/internal/token"
	"go-reloaded/internal/transform"
)

func ProcessText(in string) string {
	toks := token.Tokenize(in)

	// Numbers
	toks = transform.ApplyHex(toks)
	toks = transform.ApplyBin(toks)

	// Case tags
	toks = transform.ApplyCaseTags(toks)

	// Quotes: tighten interiors; ensure gap after closing quotes when followed by a word
	toks = transform.ApplyQuotes(toks)
	toks = transform.ApplyQuotes(toks)                // cheap second pass helps tricky adjacencies
	toks = transform.ApplySpaceAfterClosingQuote(toks)

	// Articles
	toks = transform.ApplyArticleAn(toks)

	// Normalize spaces, punctuation
	toks = transform.ApplySpaces(toks)
	toks = transform.ApplyPunctuation(toks)

	// Drop any leftover/unknown tags
	toks = transform.ApplyDropTags(toks)

	// FINAL: remove plain spaces flush against quote edges
	toks = transform.ApplyTightenQuoteEdges(toks)

	// Final sweep: collapse plain spaces and trim ends (preserve newlines)
	toks = transform.ApplySpacesWithTrim(toks, true)

	return token.Join(toks)
}