# Analysis

## 1️⃣ Problem Description (in my own words)
**go-reloaded** is a text correction tool written in Go.  
It reads an input file containing words and special “commands” such as `(hex)`, `(up, 2)`, or `(low)`.  
After identifying these commands, the program applies the corresponding transformations and writes the corrected text into a new output file.  
The goal is to understand how to handle files, parse strings, and build a modular architecture.

---

## 2️⃣ Rules (with examples)

1. **(hex)** → Converts the **previous word** from hexadecimal to decimal.  
   `1E (hex) → 30`
2. **(bin)** → Converts the **previous word** from binary to decimal.  
   `10 (bin) → 2`
3. **(up) / (low) / (cap)** → Changes the casing of the **previous word**.  
   `go (up) → GO`  
   `LOUD (low) → loud`  
   `bridge (cap) → Bridge`
4. **(up|low|cap, n)** → Applies the change to the **n previous words** (counting only words).  
   `This is so exciting (up, 2) → This is SO EXCITING`
5. **Punctuation** `. , ! ? : ;` → Attached directly to the previous word, **one space** after.  
   Grouped symbols like `...`, `!?`, `?!` remain intact.  
   `I was thinking ... You were right → I was thinking... You were right`
6. **Quotes `' … '`** → No spaces **inside** the quotes.  
   `' awesome ' → 'awesome'`, `' many words here ' → 'many words here'`
7. **`a → an`** → If the next word (ignoring spaces/commas) starts with a **vowel** or **h**.  
   `a untold → an untold`, `a, honest → an, honest`

---

## 3️⃣ Edge Cases & Decisions

- If the number is invalid: `ZZ (hex)` or `102 (bin)` → keep the word, remove the tag.  
- `(up, 0)` or incorrect syntax → safely ignored.  
- In `(up, n)`, count **only words**, not punctuation or spaces.  
- `a→an` still applies even if there is a comma in between.  
- `...` is treated as a **single punctuation mark**.

---

## 4️⃣ Pipeline vs FSM (comparison & choice)

### 🔹 What are “Stages” (Pipeline) vs “States” (FSM)
- **Stage (Pipeline)** = A **processing step**. It takes the output of the previous step, modifies it, and passes it on.  
- **State (FSM)** = A **program state**, such as "reading a word" or "reading a tag", and it changes behavior based on the input.

---

### 🔸 Pipeline Diagram (sequence of stages)

Input file  
   ↓  
[ Tokenize ]  
   ↓  
[ Numbers (hex/bin) ]  
   ↓  
[ Casing (up/low/cap) ]  
   ↓  
[ Articles (a→an) ]  
   ↓  
[ Punctuation ]  
   ↓  
[ Quotes ]  
   ↓  
Output file

---

### 🔸 FSM Diagram (flow of states)

[READ_WORD] -- '(' --> [READ_TAG] -- ')' --> [APPLY_RULE] --> back to [READ_WORD]

---

### 🧩 Criteria

| Criterion | Pipeline | FSM |
|-----------|-----------|------|
| Readability | Clean & modular | More complex (one big loop) |
| Testing | Easy (unit tests per stage) | Harder (entire flow) |
| Adding a new rule | New stage | Modify FSM logic |
| Performance | Slightly slower | Slightly faster |
| Best suited for | Data transformations | Real-time processing |

---

### 🧠 Choice
I chose the **Pipeline** approach because:
- It is **clean**, readable, and **easy to debug**.  
- Each rule is an independent, small function.  
- I can perform unit testing per stage.  
- Overall, it’s the most beginner-friendly and structured method to start with.