package token

import "fmt"

func DebugDump(toks []Tok, limit int) {
	if limit <= 0 || limit > len(toks) {
		limit = len(toks)
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%3d: %-6v %q\n", i, toks[i].K, toks[i].Text)
	}
}