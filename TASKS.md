# TASKS.md — go-reloaded

## How to run and verify
- Build: `go build .`
- Run:  `go run . <input> <output>`
- Test: `go test ./...`

## EPIC 0 — CLI & wiring
- [x] 0.1 Args: if not exactly 2, print `Usage: goreloaded <input> <output>` and exit.
- [x] 0.2 Read input file. If missing/unreadable → print error and exit.
- [x] 0.3 If output exists, ask: `Overwrite? (y/n):` → on `y` write, else cancel.
- [x] 0.4 Call `pipeline.ProcessText` and write the result.

**Accept when:** manual run works; wrong args show usage; overwrite prompt appears; file is written.

## EPIC 1 — Tokenization
- [x] 1.1 Tokenize into: Group(`...`,`!?`,`?!`), Punct(`.,!?:;`), Quote(`'`), Tag(`(low, 2)` etc.), Space, Word
- [x] 1.2 Join returns the exact concatenation of tokens (no implicit trimming)

**Accept when:** tokens split `Punctuation tests are ... kinda boring ,what do you think ?` the way rules require; Join reproduces original string.

## EPIC 2 — Transforms
- [x] 2.1 `(hex)`: convert previous Word from hex→decimal; drop tag; if invalid keep word/drop tag; if no previous word drop tag.
- [x] 2.2 `(bin)`: same pattern as hex with base 2.
- [x] 2.3 `(up|low|cap[, n])`: apply to previous n **Word** tokens; `(cap)` uses Title Case (first letter upper, rest lower).
- [x] 2.4 Quotes: remove spaces just inside paired `'` … `'` for single words or spans.
- [x] 2.5 Article: `a`→`an` if next Word starts with vowel or `h` (case-insensitive). Preserve capital "A"→"An" when needed.
- [x] 2.6 Punctuation: remove spaces **before** punct (unless newline), ensure **one space after**. Groups stay tight.
- [x] 2.7 Space normalization: collapse multiple spaces while preserving newlines.

**Accept when:** tiny manual strings behave exactly like examples.

## EPIC 3 — Golden tests
- [x] 3.1 Add golden pairs with descriptive names:
  - `case_transforms.txt` / `case_transforms.want.txt`
  - `number_conversion.txt` / `number_conversion.want.txt`
  - `article_correction.txt` / `article_correction.want.txt`
- [x] 3.2 `internal_test/golden_test.go` loops `testdata/*.txt` and compares to `*.want.txt`.

**Accept when:** `go test ./...` passes; outputs match expected text exactly.

## EPIC 4 — Polish
- [x] 4.1 README: usage, examples, test instructions.
- [x] 4.2 AGENTS.md: already present; skim for accuracy.
- [x] 4.3 Added comprehensive documentation: analysis.md, final_report.md.
- [x] 4.4 MIT License: proper LICENSE.txt file.
- [x] 4.5 Code quality: `go fmt`, `go vet` clean.

**Accept when:** repo is easy to read, run, and grade.

## ✅ PROJECT COMPLETE

All EPICs completed successfully:
- ✅ **100% test pass rate**
- ✅ **Production-ready code**
- ✅ **Professional documentation**
- ✅ **Clean architecture**