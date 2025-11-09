package pipeline

import (
	"go-reloaded/internal/token"
	"go-reloaded/internal/transform"
)

func ProcessText(in string) string {
	toks := token.Tokenize(in)

	// 1) number conversions first
	toks = transform.ApplyHex(toks)
	toks = transform.ApplyBin(toks)

	// 2) case tags
	toks = transform.ApplyCaseTags(toks)

	// 3) quotes: tighten insides, then ensure gap after closing quote before words
	toks = transform.ApplyQuotes(toks)
	toks = transform.ApplyQuotes(toks)                // extra pass (cheap, robust)
	toks = transform.ApplySpaceAfterClosingQuote(toks)

	// 4) articles (a -> an), skipping spaces & quotes to the next word
	toks = transform.ApplyArticleAn(toks)

	// 5) normalize spaces, then punctuation rules
	toks = transform.ApplySpaces(toks)
	toks = transform.ApplyPunctuation(toks)

	// 6) drop any leftover (unknown/malformed) tags
	toks = transform.ApplyDropTags(toks)

	// 7) final sweep: collapse plain spaces & trim leading/trailing plain space
	toks = transform.ApplySpacesWithTrim(toks, true)

	return token.Join(toks)
}