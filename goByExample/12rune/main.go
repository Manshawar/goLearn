package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("s:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("rune Count", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("(%d %x)", idx, runeValue)
	}
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("(%d %c %d)", i, runeValue, width)
		w = width
		examineRune(runeValue)
	}
}
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	} else {
		fmt.Println("not found")
	}
}
