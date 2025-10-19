# go-reloaded   
### Week 1 – Documentation Phase

---

## 📘 Project Description
This project is about creating a **text reloader tool** in Go.  
It reads a text file as input, detects transformation instructions like `(hex)`, `(up, 2)`, `(low)` etc.,  
applies the corresponding rules, and writes a corrected version of the text into a new file.

During this **first week**, no Go code is written yet.  
The focus is only on **understanding the problem**, **designing the logic**, and **defining success test cases**.

---

## 📂 Folder Structure

go-reloaded/
│
├── docs/
│ ├── analysis.md # Detailed description of the problem, rules, and chosen method (Pipeline)
│ └── golden-tests.md # The “Golden Test Set” defining all success cases
│
└── README.md


---

## 🧩 Summary of What Has Been Done

✅ Complete problem analysis written in `docs/analysis.md`  
✅ Comparison between **Pipeline** and **FSM** methods  
✅ Final decision: **Pipeline method** (modular, clear, auditable)  
✅ All success test cases written in `docs/golden-tests.md`

---

### Author
**Theodore Konstadiniakos**  
Zone01 Athens · 2025
