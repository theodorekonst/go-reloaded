# go-reloaded

A complete text transformation tool written in Go that processes files with special transformation tags and outputs corrected text. The project follows clean architecture principles with full separation of concerns.

## 🚀 Quick Start

```
# Build the program
go build .

# Run with input and output files
go run . input.txt output.txt

# Run tests
go test ./...
```

## ✨ What It Does

Transform text using special tags:

```
Input:  it (cap) was 42 (hex) and a honest (up, 2) mistake
Output: It was 66 and AN HONEST mistake
```

## 🔧 Transformation Rules

| Rule | Example | Result |
|------|---------|--------|
| `(hex)` | `42 (hex)` | `66` |
| `(bin)` | `10 (bin)` | `2` |
| `(up)` | `word (up)` | `WORD` |
| `(low)` | `WORD (low)` | `word` |
| `(cap)` | `word (cap)` | `Word` |
| `(up, n)` | `these words (up, 2)` | `THESE WORDS` |
| Article | `a apple` | `an apple` |
| Punctuation | `word ,space` | `word, space` |
| Quotes | `' spaced '` | `'spaced'` |

## 📁 Project Structure

```
go-reloaded/
├── main.go                    # CLI entry point
├── internal/
│   ├── io/                   # File operations
│   ├── token/                # Text tokenization
│   ├── transform/            # Transformation rules
│   └── pipeline/             # Processing pipeline
├── testdata/                 # Test files
└── internal_test/            # Test suite
```

## 🧪 Testing

```
# Run all tests
go test ./...

# Test with provided examples
go run . testdata/case_transforms.txt output.txt
go run . testdata/number_conversion.txt output.txt
go run . testdata/article_correction.txt output.txt

# Create your own test file
# Create input.txt with: hello (cap) world
go run . input.txt result.txt
# Check result.txt - should contain: Hello world
```

## 💡 Usage Examples

### Number Conversions

Create a file `numbers.txt` with:

```
Value: 2A (hex) and 1010 (bin)
```

Run:

```
go run . numbers.txt output.txt
```

Result:

```
Value: 42 and 10
```

### Case Transformations

Create a file `cases.txt` with:

```
make this (up, 3) text better
```

Run:

```
go run . cases.txt output.txt
```

Result:

```
MAKE THIS TEXT better
```

### Article Corrections

Create a file `articles.txt` with:

```
It was a honest mistake and a apple
```

Run:

```
go run . articles.txt output.txt
```

Result:

```
It was an honest mistake and an apple
```

## 🏗️ Architecture

Pipeline Design for clean separation of concerns:

1. **Tokenize** → Parse text into tokens
2. **Hex** → Convert hexadecimal numbers
3. **Bin** → Convert binary numbers
4. **Case** → Apply case transformations
5. **Quotes** → Tighten quoted text
6. **Article** → Fix a→an corrections
7. **Spaces** → Normalize spacing
8. **Punctuation** → Format spacing
9. **Join** → Reconstruct text

## ✅ Features

* ✅ Hex/binary to decimal conversion
* ✅ Case transformations (up/low/cap with ranges)
* ✅ Smart article correction (a→an)
* ✅ Punctuation spacing rules
* ✅ Quote tightening
* ✅ Error handling for invalid inputs
* ✅ Comprehensive test suite
* ✅ Clean CLI interface

## 🎯 Requirements

* Go 1.25
* No external dependencies

## 📝 License

MIT License - see [LICENSE.txt](LICENSE.txt) for details.

This project is part of a coding exercise demonstrating clean Go architecture and text processing techniques.