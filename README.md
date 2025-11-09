# Go-reloaded

**Version:** 1.0.0

A complete text transformation tool written in Go that changes text files using special commands. The project follows clean code design where each part has its own job.

## 🚀 Quick Start

```
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
├── main.go                    # Command line program
├── internal/
│   ├── io/                   # Reading and writing files
│   ├── token/                # Breaking text into pieces
│   ├── transform/            # Text change rules (14 files)
│   └── pipeline/             # Controls order of changes
├── testdata/                 # Example test files
└── internal_test/            # Test programs
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

## 🏗️ How It Works

Pipeline Design where each step does one job:

1. **Break Apart** → Split text into pieces
2. **Convert Hex** → Change hex numbers to regular numbers
3. **Convert Binary** → Change binary numbers to regular numbers
4. **Change Cases** → Make text uppercase, lowercase, or title case
5. **Fix Quotes** → Remove extra spaces inside quotes
6. **Fix Articles** → Change "a" to "an" when needed
7. **Fix Spaces** → Remove extra spaces
8. **Fix Punctuation** → Put punctuation in the right place
9. **Put Together** → Combine everything back into text

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

This project shows clean Go code design and text processing techniques that are easy to understand and learn from.