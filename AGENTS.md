# TDD Backlog for go-reloaded

---

## EPIC 0 — Project & Scaffolding

### Task 0.1 — Repo bootstrap & test harness

**Functionality:** Compile a minimal project and run tests end-to-end.

**TDD:** Add `internal/pipeline/process_test.go` with a placeholder failing test to ensure the harness is wired.

**Implementation:** Init `go mod`; folders `cmd/goreloaded`, `internal/{io,token,transform,pipeline}`; Makefile with test target.

**Validation:** `go test ./...` passes (placeholder test now green).

---

### Task 0.2 — CLI argument handling

**Functionality:** `goreloaded <input> <output>` with helpful error on missing args.

**TDD:** Test main invoked without args exits non-zero and prints usage; with args, it does not error (processing stubbed).

**Implementation:** `cmd/goreloaded/main.go` parses args and calls `ProcessText`.

**Validation:** Tests pass on both failure and success paths.

---

## EPIC 1 — File I/O

### Task 1.1 — Read file

**Functionality:** Read entire input file as string.

**TDD:** Test reading existing fixture vs nonexistent path (expect error).

**Implementation:** `internal/io/read.go`: `ReadFile(path string) (string, error)`.

**Validation:** Unit tests green.

---

### Task 1.2 — Write file

**Functionality:** Write processed text to output file.

**TDD:** Round-trip: write text → read back equals; error on unwritable path.

**Implementation:** `internal/io/write.go`: `WriteFile(path, s string) error`.

**Validation:** Unit tests green.

---

## EPIC 2 — Tokenization (non-destructive scan)

### Task 2.1 — Token model & kinds

**Functionality:** Represent tokens as Word, Tag, Punct (single & grouped), Quote.

**TDD:** Table-driven: classify `.`, `,`, `!`, `?`, `:`, `;`, grouped `...`, `!?`, `?!`; tags like `(hex)`, `(bin)`, `(up)`, `(cap, 3)`, raw quote `'`. Rules for grouping and tag shapes come from your **analysis**.

**Implementation:** `internal/token/token.go` with
```go
type Token struct { Value string; Kind TokenKind }
```

**Validation:** Classification tests pass.

---

### Task 2.2 — Tokenizer

**Functionality:** Produce an ordered token stream preserving groups and parentheses.

**TDD:** `"Hello (up) world ... ok?!"` → tokens reflect `(up)`, `...`, `?!` as single items; words preserved.

**Implementation:** `internal/tokenize/tokenize.go`: `Tokenize(input string) []Token`.

**Validation:** Token sequences match expectations.

---

## EPIC 3 — Transformations (Pipeline Stages)

Apply stages in this order: Numbers → Case → Articles → Punctuation → Quotes (after Tokenize, before Join).
This order is explicitly defined in your analysis and reflected in the mixed golden test set.

### Task 3.1 — Numbers: (hex) + (bin)

**Functionality:** Replace previous word with decimal value for valid hex/bin; otherwise keep word and drop tag.

**TDD:**
- `"1E (hex)"` → `"30"`, `"10 (bin)"` → `"2"`.
- Invalids remain: `"ZZ (hex)"` → `"ZZ"`, `"102 (bin)"` → `"102"`.
- (From rules and golden test set.)

**Implementation:** `internal/transform/number.go`: `TransformHexBin([]Token) []Token` using `strconv.ParseInt` with bases 16/2.

**Validation:** Unit tests green.

---

### Task 3.2 — Case: (up), (low), (cap) and ranged (…, n)

**Functionality:** Change case of the previous word(s); `(n)` counts only word tokens (skip spaces/punct).

**TDD:**
- Singles: `go (up)` → `GO`, `LOUD (low)` → `loud`, `bridge (cap)` → `Bridge`.
- Ranges: `"This is so exciting (up, 2)"` → `"This is SO EXCITING"`.
- Ignore malformed: `(up, 0)` or bad syntax = no change.
- (All per analysis + golden test set.)

**Implementation:** `internal/transform/case.go` with helper to walk back n word tokens and mutate values.

**Validation:** Table tests (including cap chains like "It Was The Age…") pass.

---

### Task 3.3 — Articles: a → an

**Functionality:** Convert `a` to `an` if the next word, ignoring intervening spaces/commas, starts with a vowel or h.

**TDD:**
- `"a apple"` → `"an apple"`, `"a, honest"` → `"an, honest"`, `"a car"` → `"a car"`.
- Paragraph cases (e.g., "an unusual…", "an hour…") from golden test set.

**Implementation:**
- `internal/transform/article.go`: `TransformArticles([]Token) []Token`
- `startsWithVowelOrH(string) bool`, scanner that skips commas/spaces.

**Validation:** Unit tests and paragraph samples pass.

---

### Task 3.4 — Punctuation: spacing & grouping

**Functionality:** Attach punctuation to the left, ensure one space to the right; keep groups `...`, `!?`, `?!` intact.

**TDD:**
- `"Punctuation tests are ... kinda boring ,what do you think ?"` → `"Punctuation tests are... kinda boring, what do you think?"`
- `"Wait ...what?! are you sure !?"` → `"Wait... what?! are you sure!?"`
- (from golden test set)

**Implementation:** `internal/transform/punct.go` enforces canonical spacing without breaking grouped tokens.

**Validation:** Table tests pass.

---

### Task 3.5 — Quotes: trim inner spaces

**Functionality:** Remove spaces inside single quotes only, keep quotes in place.

**TDD:**
- `"' awesome '"` → `"'awesome'"`
- `"' many words here '"` → `"'many words here'"`
- (from analysis and golden test set)

**Implementation:** `internal/transform/quotes.go` collapses space tokens at quote boundaries within a quoted span.

**Validation:** Unit tests pass.

---

## EPIC 4 — Orchestration & Joining

### Task 4.1 — Orchestrate the pipeline

**Functionality:** `ProcessText` executes: Tokenize → Numbers → Case → Articles → Punctuation → Quotes → Join.

**TDD:** Integration test with mixed case: `"This is 1E (hex) awesome (up)"` → `"This is 30 AWESOME"`.
(Order validated per analysis.)

**Implementation:** `internal/pipeline/process.go`: `ProcessText(string) string`.

**Validation:** Integration test green.

---

### Task 4.2 — Join tokens to string

**Functionality:** Convert tokens back to a canonical string with spacing rules preserved.

**TDD:** Joining round-trips: token stream → string equals expected for representative samples (including punctuation groups and quotes).

**Implementation:** `internal/pipeline/join.go`: `Join([]Token) string`.

**Validation:** Join tests pass.

---

### Task 4.3 — CLI end-to-end

**Functionality:** Wire `ProcessText` to CLI: read input, process, write output.

**TDD:** E2E test uses temp files; expected equals actual for a small scenario.

**Implementation:** Complete `main.go` flow with error handling.

**Validation:** E2E test passes.

---

## EPIC 5 — Golden Suites

### Task 5.1 — Golden: Core

**Functionality:** Encode Core examples as table tests.

**TDD:** Inputs/outputs exactly as the "Core Cases" (cap/up/low/ranges; hex+bin; a→an; punctuation spacing) from golden test set.

**Implementation:** `test/golden_core_test.go` drives `ProcessText` and compares strings.

**Validation:** All Core tests green.

---

### Task 5.2 — Golden: Tricky

**Functionality:** Encode Tricky examples (invalid numbers, ranges over punctuation, quotes, comma with article, mixed tags, ellipsis + interrobang).

**TDD:** Inputs/outputs exactly as listed in golden test set.

**Implementation:** `test/golden_tricky_test.go`.

**Validation:** All Tricky tests green.

---

### Task 5.3 — Golden: Big Paragraph

**Functionality:** Full combined-rule test for real-world paragraph.

**TDD:** Use the long input → expect the exact long output in the golden test set.

**Implementation:** `test/golden_paragraph_test.go`.

**Validation:** Paragraph test green (proves ordering/interplay correctness).

---

## EPIC 6 — Robustness & Quality

### Task 6.1 — Malformed tags & no-op safety

**Functionality:** Gracefully ignore `(up, 0)`, `(cap,)`, `(low, -1)`, etc.

**TDD:** Table tests ensure no panics and no unintended mutations (per Edge Cases in analysis).

**Implementation:** Defensive parsers with defaults and early returns.

**Validation:** Tests pass.

---

### Task 6.2 — Idempotence & stability

**Functionality:** Re-processing an already processed string does not change it again (esp. punctuation/quotes).

**TDD:** Property-style test: `ProcessText(x) == ProcessText(ProcessText(x))`.

**Implementation:** Ensure joiner and spacing rules are canonical.

**Validation:** Property tests green.

---

### Task 6.3 — Race-safety & perf sanity (optional)

**Functionality:** Safe for concurrent execution (no shared global state).

**TDD:** `go test -race`; micro-bench for large input (no strict SLA, just baseline).

**Implementation:** Keep functions pure; avoid shared mutables.

**Validation:** `-race` clean; benchmark runs.

---

## EPIC 7 — Dev Experience

### Task 7.1 — README & runnable examples

**Functionality:** Clear usage docs with copy-paste example matching golden expectations.

**TDD:** CI script runs the example and diffs with expected output.

**Implementation:** `README.md` shows `go run ./cmd/goreloaded in.txt out.txt`.

**Validation:** CI doc-example check passes.

---

### Task 7.2 — Lint/Vet gates

**Functionality:** Basic static checks.

**TDD:** CI fails on `go vet` issues or formatting drift.

**Implementation:** CI workflow with `go vet ./...` and a fmt check.

**Validation:** Pipeline green.

---

## How AI Agents Help (per task)

- **Before coding:** Propose table-driven tests and corner cases extracted from the analysis and golden test set.
- **During TDD:** Suggest minimal diffs to satisfy the current failing test only (no over-implementation).
- **After green:** Recommend small refactors (helpers, names) with no behavior change.
- **Review checklist:** Ensure each PR references the relevant rule paragraph and golden example implemented.

---

## Definition of Done

- All Golden suites (Core, Tricky, Big Paragraph) pass exactly as specified in the golden test set.
- All rule semantics and edge-case decisions are implemented per analysis.
- CLI produces correct output files for provided inputs; tests and CI are green.