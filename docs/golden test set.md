# Golden Test Set (Success Cases)

This file describes the expected behavior of the **go-reloaded** tool.  
Each test includes an **Input** and its corresponding **Expected Output**.

---

## A. Core Cases (from the subject)

### 1) Cap / Up / Low / Ranges
**Input**  
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

**Expected Output**  
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

---

### 2) Hex + Bin
**Input**  
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

**Expected Output**  
Simply add 66 and 2 and you will see the result is 68.

---

### 3) Article a→an
**Input**  
There is no greater agony than bearing a untold story inside you.

**Expected Output**  
There is no greater agony than bearing an untold story inside you.

---

### 4) Punctuation spacing + groups
**Input**  
Punctuation tests are ... kinda boring ,what do you think ?

**Expected Output**  
Punctuation tests are... kinda boring, what do you think?

---

## B. Tricky Cases

### 5) Invalid hex/bin
**Input**  
Values: ZZ (hex) and 102 (bin) should remain as-is.

**Expected Output**  
Values: ZZ and 102 should remain as-is.

---

### 6) Ranges over punctuation
**Input**  
smart, quick fox (cap, 3) jumps! really FAST (low, 1).

**Expected Output**  
Smart, Quick Fox jumps! really fast.

---

### 7) Quotes with inner spaces
**Input**  
As they said: ' awesome ' and ' many words here ' are examples.

**Expected Output**  
As they said: 'awesome' and 'many words here' are examples.

---

### 8) Article with comma in between
**Input**  
It was a, honest mistake and a unusual choice.

**Expected Output**  
It was an, honest mistake and an unusual choice.

---

### 9) Mixed tags & order of application
**Input**  
This is fine (up) and also great (low, 2) TODAY.

**Expected Output**  
THIS is fine and also great today.

---

### 10) Ellipsis + Interrobang
**Input**  
Wait ...what?! are you sure !?

**Expected Output**  
Wait... what?! are you sure!?

---

## C. Big Paragraph (combined rules)
**Input**  
it (cap) was a honest idea, but the plan had 1E (hex) steps, and only 10 (bin) were done ...why ? ' maybe later ' she said. the quick fix (up, 3) was proposed, yet a unusual error appeared; we tried again (low, 2) AFTER THAT, and a hour passed! finally, we wrote the Brooklyn bridge (cap) story, and it was great (cap, 4) , right ?

**Expected Output**  
It was an honest idea, but the plan had 30 steps, and only 2 were done... why? 'maybe later' she said. The QUICK FIX was proposed, yet an unusual error appeared; we tried again after that, and an hour passed! Finally, we wrote the Brooklyn Bridge story, and It Was Great, right?