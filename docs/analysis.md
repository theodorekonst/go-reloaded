# Analysis

## 1️⃣ Problem Description

go-reloaded is a text transformation tool written in Go that changes text files using special commands. It reads a text file, finds commands like `(hex)`, `(up, 2)`, or `(cap)`, makes the changes, and saves the result to a new file.

The goal is to build a strong, easy-to-maintain program that can handle many different text changes while keeping the original text format.

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
- Each step does one job and passes the result to the next step
- Like a factory: Break text apart → Convert numbers → Change cases → Fix spacing → Done

**State Machine**
- One program that changes what it does based on what it finds
- Like a smart robot: Reading text → Found a command → Apply the rule → Keep going

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

### How We Handle Errors

- **Handle errors well:** When input is wrong → keep the text, remove the command
- **Keep format safe:** Never break the original text layout
- **Stay safe:** Unknown commands are ignored and passed to other steps

## 6️⃣ Testing Strategy

### Golden Tests

- **Clear names:** `case_transforms.txt`, `number_conversion.txt`, `article_correction.txt`, `final_all_cases.txt`
- **Input/Output pairs:** Each `.txt` file has a matching `.want.txt` file with expected results
- **Tests everything:** All transformation rules are tested
- **Full testing:** The entire pipeline is tested together

### Test Categories

- **case_transforms** → Tests all case operations and ranges
- **number_conversion** → Tests hex/bin conversions and error cases
- **article_correction** → Tests a→an grammar rules

## 7️⃣ Project Structure

```
go-reloaded/
├── main.go                    # Command line program
├── internal/
│   ├── io/file.go            # Reading and writing files
│   ├── token/token.go        # Breaking text into pieces and putting it back
│   ├── transform/            # 14 files, each handles one type of change
│   │   ├── convert.go        # Changes hex/binary numbers
│   │   ├── case.go           # Changes uppercase/lowercase
│   │   ├── article.go        # Fixes "a" vs "an"
│   │   ├── quotes.go         # Fixes quote spacing
│   │   ├── punct.go          # Fixes punctuation spacing
│   │   ├── space.go          # Fixes extra spaces
│   │   └── ... (8 more)      # Other transformation rules
│   └── pipeline/pipeline.go  # Controls the order of changes
├── testdata/                 # Test files with examples
└── internal_test/            # Test runner program
```

### Design Rules

- **One job per file:** Each file does one thing well
- **Keep things separate:** Command line, text processing, and file handling are in different places
- **No extra libraries:** Uses only built-in Go features
- **Easy to learn:** Simple, clear code that new developers can understand

## 8️⃣ Success Criteria

### What the Program Must Do

- ✅ All text transformation rules work correctly
- ✅ Weird cases are handled well
- ✅ Original text format is kept safe
- ✅ Errors are handled properly

### Code Quality Requirements

- ✅ Clean, easy-to-maintain code structure
- ✅ 100% test coverage of all transformation rules
- ✅ Works on Windows, Mac, and Linux
- ✅ Professional command line interface

### Speed Requirements

- ✅ Processing time grows linearly with file size
- ✅ Uses memory efficiently
- ✅ Fast processing for normal text files