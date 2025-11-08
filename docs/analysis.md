# Analysis

## 1️⃣ Problem Description

go-reloaded is a text transformation tool written in Go that processes files with special transformation tags and outputs corrected text. It reads an input file, identifies transformation commands like `(hex)`, `(up, 2)`, or `(cap)`, applies the corresponding changes, and writes the result to an output file.

The goal is to build a robust, maintainable system that handles multiple transformation types while preserving text structure and following clean architecture principles.

## 2️⃣ Transformation Rules

### Number Conversions

**`(hex)`** → Converts the previous word from hexadecimal to decimal
```
42 (hex) → 66
```

**`(bin)`** → Converts the previous word from binary to decimal
```
10 (bin) → 2
```

### Case Transformations

**`(up)`** → Makes the previous word UPPERCASE
```
word (up) → WORD
```

**`(low)`** → Makes the previous word lowercase
```
WORD (low) → word
```

**`(cap)`** → Makes the previous word Title Case
```
word (cap) → Word
```

**`(up|low|cap, n)`** → Applies transformation to n previous words (counting only words)
```
these words (up, 2) → THESE WORDS
```

### Text Corrections

**Article correction** → Changes a to an before vowels or h
```
a apple → an apple, a honest → an honest
```

**Punctuation spacing** → Attaches punctuation to previous word, one space after
```
word ,space → word, space
```

**Punctuation groups** → Treats ..., !?, ?! as single units
```
word ... space → word... space
```

**Quote tightening** → Removes spaces inside quotes
```
' spaced words ' → 'spaced words'
```

## 3️⃣ Edge Cases & Error Handling

- **Invalid numbers:** `ZZ (hex)` or `102 (bin)` → keep word, remove tag
- **Missing previous word:** `(cap) hello` → remove tag, keep text
- **Invalid syntax:** `(up, 0)` or malformed tags → safely ignored
- **Word counting:** In `(up, n)`, count only words, skip punctuation/spaces
- **Article with punctuation:** `a, honest` → `an, honest` (works across punctuation)
- **Preserve structure:** Maintain original line breaks and formatting

## 4️⃣ Architecture Decision: Pipeline vs FSM

### What are the approaches?

**Pipeline (Assembly Line)**
- Each stage does one specific job and passes result to next stage
- Like a factory: Tokenize → Convert → Transform → Format → Output

**FSM (State Machine)**
- Single processor that changes behavior based on current state
- Like a smart robot: Reading → Found Tag → Apply Rule → Continue

### Pipeline Implementation

```
Input Text
    ↓
[ Tokenize ] ← Split into Word, Space, Punct, Tag tokens
    ↓
[ Hex Conversion ] ← Handle (hex) tags
    ↓
[ Bin Conversion ] ← Handle (bin) tags  
    ↓
[ Case Transforms ] ← Handle (up), (low), (cap) tags
    ↓
[ Quote Tightening ] ← Remove spaces in quotes
    ↓
[ Article Correction ] ← Fix a→an
    ↓
[ Space Normalization ] ← Collapse multiple spaces
    ↓
[ Punctuation Spacing ] ← Fix punctuation attachment
    ↓
Output Text
```

### Why Pipeline Order Matters

1. **Numbers first** → Independent of other transforms
2. **Case after numbers** → Might need to transform converted numbers
3. **Quotes before articles** → Article rules might be affected by quote changes
4. **Articles before spacing** → Grammar fixes before formatting
5. **Space normalization before punctuation** → Clean up spaces first
6. **Punctuation last** → Final formatting pass

### Architecture Comparison

| Criterion | Pipeline | FSM |
|-----------|----------|-----|
| Readability | ✅ Clean & modular | ❌ Complex single function |
| Testing | ✅ Unit test each stage | ❌ Must test entire flow |
| Debugging | ✅ Easy to isolate issues | ❌ Hard to find problems |
| Maintenance | ✅ Modify one stage | ❌ Change affects whole system |
| Team Development | ✅ Parallel work on stages | ❌ Single point of conflict |
| Performance | ❌ Multiple passes | ✅ Single pass |
| Memory Usage | ❌ Intermediate results | ✅ Process as you go |

### 🎯 Decision: Pipeline Architecture

Chosen Pipeline because:

- **Code Quality:** More readable and maintainable
- **Testing:** Each transform independently testable
- **Team Friendly:** Multiple developers can work on different stages
- **Debugging:** Easy to trace issues through pipeline
- **Extensibility:** New rules are just new pipeline stages
- **Learning:** Better for junior developers to understand

## 5️⃣ Implementation Strategy

### Token-Based Processing

```go
type Token struct {
    K    TokenType  // Word, Space, Punct, Tag, Group, Quote
    Text string     // Actual content
}
```

**Benefits:**
- Preserves original spacing and structure
- Enables precise transformations without losing context
- Maintains exact formatting including line breaks

### Transform Functions

Each transform follows the same pattern:

```go
func ApplyTransform(tokens []Token) []Token {
    // 1. Find relevant tokens
    // 2. Apply transformation
    // 3. Remove processed tags
    // 4. Return modified token stream
}
```

### Error Handling Philosophy

- **Graceful degradation:** Invalid input → keep original, drop tag
- **Preserve structure:** Never corrupt original text format
- **Fail safe:** Unknown tags pass through to other transforms

## 6️⃣ Testing Strategy

### Golden Tests

- **Descriptive names:** `case_transforms.txt`, `number_conversion.txt`, `article_correction.txt`
- **Input/Output pairs:** Each `.txt` has corresponding `.want.txt`
- **Comprehensive coverage:** All transformation rules tested
- **Integration testing:** Full pipeline validation

### Test Categories

- **case_transforms** → Tests all case operations and ranges
- **number_conversion** → Tests hex/bin conversions and error cases
- **article_correction** → Tests a→an grammar rules

## 7️⃣ Project Structure

```
go-reloaded/
├── main.go                    # CLI interface
├── internal/
│   ├── io/file.go            # File operations & overwrite handling
│   ├── token/token.go        # Tokenization & reconstruction
│   ├── transform/            # One file per transformation rule
│   │   ├── convert.go        # Hex/binary conversions
│   │   ├── case.go           # Case transformations
│   │   ├── article.go        # Article corrections
│   │   ├── quotes.go         # Quote tightening
│   │   ├── punct.go          # Punctuation spacing
│   │   └── space.go          # Space normalization
│   └── pipeline/pipeline.go  # Transform orchestration
├── testdata/                 # Golden test files
└── internal_test/            # Test runner
```

### Design Principles

- **Single Responsibility:** Each file has one clear purpose
- **Separation of Concerns:** CLI, processing, and I/O are separate
- **No External Dependencies:** Pure Go standard library
- **Junior Developer Friendly:** Simple, readable code structure

## 8️⃣ Success Criteria

### Functional Requirements

- ✅ All transformation rules implemented correctly
- ✅ Edge cases handled gracefully
- ✅ Original text structure preserved
- ✅ Comprehensive error handling

### Quality Requirements

- ✅ Clean, maintainable architecture
- ✅ 100% test coverage of transformation rules
- ✅ Cross-platform compatibility
- ✅ Professional CLI interface

### Performance Requirements

- ✅ Linear time complexity O(n)
- ✅ Reasonable memory usage
- ✅ Fast execution for typical text files