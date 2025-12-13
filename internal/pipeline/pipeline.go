package pipeline

import (
	"go-reloaded/internal/token"
	"go-reloaded/internal/transform"
)

func ProcessText(in string) string {
	toks := token.Tokenize(in)

	// Validate tags (must have space before)
	toks = transform.ValidateTags(toks)

	// Numbers
	toks = transform.ApplyHex(toks)
	toks = transform.ApplyBin(toks)

	// Case tags
	toks = transform.ApplyCaseTags(toks)

	// Articles (AFTER case transforms)
	toks = transform.ApplyArticleAn(toks)

	// QUOTES
	toks = transform.ApplyQuotes(toks)
	toks = transform.ApplyQuotes(toks) // cheap second pass for tricky adjacencies
	toks = transform.ApplyQuoteSpacingFix(toks) // Comprehensive quote spacing fix
	toks = transform.ApplyApostropheSpacing(toks) // Fix apostrophe spacing
	toks = transform.ApplySpaceAfterClosingQuote(toks)
	toks = transform.ApplySpaceBeforeOpeningQuote(toks) // NEW

	// Normalize spaces, punctuation
	toks = transform.ApplySpaces(toks)
	toks = transform.ApplyPunctuation(toks)

	// CLEANUP / SPECIALS
	toks = transform.ApplyDashQuoteTight(toks) // tighten —'quote' (remove space)
	toks = transform.ApplyCaseNextMarker(toks) // transform next word after case-range tag
	toks = transform.CapitalizeI(toks)         // capitalize personal pronoun "I"
	toks = transform.ApplyDropTags(toks)

	// FINAL: remove plain spaces flush against quote edges
	toks = transform.ApplyTightenQuoteEdges(toks)
	// Apply quotes multiple times to ensure all spacing is fixed
	toks = transform.ApplyQuotes(toks)
	toks = transform.ApplyQuotes(toks)
	toks = transform.ApplyQuoteSpacingFix(toks)

	// Final sweep: collapse plain spaces and trim ends (preserve newlines)
	toks = transform.ApplySpacesWithTrim(toks, true)
	// One more punctuation pass to fix any remaining spacing issues
	toks = transform.ApplyPunctuation(toks)
	// Final space cleanup
	toks = transform.ApplySpacesWithTrim(toks, true)
	// Final spacing fix for all edge cases - MUST BE LAST
	toks = transform.ApplyFinalSpacingFix(toks)
	// One more space cleanup after final fix
	toks = transform.ApplySpacesWithTrim(toks, true)

	return token.Join(toks)
}
