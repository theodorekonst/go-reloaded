# AGENTS.md — go-reloaded

## Goal (what this tool does)
Command-line app:

```
go run . <input> <output>
```

Read the input file, apply the rules, write the output file. Follow these rules:

- `(hex)`: previous word is hex → replace with decimal (drop the tag)
- `(bin)`: previous word is binary → replace with decimal (drop the tag)
- `(up)`, `(low)`, `(cap)`: change the previous word's case
- `(up, n)`, `(low, n)`, `(cap, n)`: apply to previous **n words** (count words only, skip punctuation)
- `(cap)` = **Title Case** (first letter uppercase, rest lowercase)
- Punctuation `.,!?:;` hugs the previous word; ensure **one space after**
- Punctuation groups `...`, `!?`, `?!` stay **tight** to previous word; one space after
- Quotes `' ... '` → remove spaces inside → `'...'` (works for single word or spans)
- `a` → `an` if next word starts with a vowel or **h** (case-insensitive)
- If a tag has no previous word (e.g., at file start), **drop the tag**
- If `(hex)`/`(bin)` number is invalid (e.g., `0x1E`), **keep the word**, **drop the tag**
- Preserve any existing line breaks

## Architecture (folders)
- `main.go` (root): CLI. Validates args, asks overwrite (y/n) if needed, reads input, calls pipeline, writes output.
- `internal/token`: `Tokenize(string) []Tok` and `Join([]Tok) string`. Token kinds: Word, Space, Quote, Punct, Group, Tag.
- `internal/transform`: one file per rule:
  - `convert.go`: `(hex)`, `(bin)`
  - `case.go`: `(up|low|cap[, n])`
  - `quotes.go`: tighten quotes
  - `article.go`: `a` → `an`
  - `space.go`: normalize spacing
  - `punct.go`: punctuation spacing and groups
- `internal/pipeline`: `ProcessText(string) string` → calls transforms in this order:
  1. Tokenize
  2. Hex
  3. Bin
  4. Case (up/low/cap[, n])
  5. Quotes (tighten)
  6. Article (a → an)
  7. Space Normalization
  8. Punctuation (attach + spacing)
  9. Join

## Usage

```
go run . <input> <output>
```

If `<output>` exists: prompt `Overwrite? (y/n):`.

## Testing
- **Golden tests** live in `testdata/` as pairs: `X.txt` → `X.want.txt`
- Test runner: `go test ./...`
- A test compares `ProcessText(X.txt)` vs `X.want.txt`. If different → FAIL with a diff-like message.

## Done definition
- All transforms implemented per rules above
- Golden tests pass (`go test ./...`)
- Running the subject examples yields the exact expected output