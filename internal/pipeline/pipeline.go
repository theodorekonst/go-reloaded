package pipeline

import (
	"go-reloaded/internal/token"
	"go-reloaded/internal/transform"
)

func ProcessText(in string) string {
	toks := token.Tokenize(in)

	toks = transform.ApplyHex(toks)
	toks = transform.ApplyBin(toks)

	toks = transform.ApplyCaseTags(toks)

	toks = transform.ApplyQuotes(toks)                 // tighten inside
	toks = transform.ApplySpaceAfterClosingQuote(toks) // add gap before next Word

	toks = transform.ApplyArticleAn(toks)

	toks = transform.ApplySpaces(toks)                 // collapse plain spaces (preserve newlines)
	toks = transform.ApplyPunctuation(toks)            // hug left, 1 space after
	toks = transform.ApplyDropTags(toks)               // remove any leftover ( ... )
	toks = transform.ApplySpacesWithTrim(toks, true)   // final sweep: collapse & trim ends

	return token.Join(toks)
}