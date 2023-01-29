package examples

import (
	"fmt"
	"unicode/utf8"
)

func runeCount(s string) {
	fmt.Println(s, " has length of ", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\n", s[i])
	}
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d \n", runeValue, idx)
	}
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == '🫀' {
		fmt.Println("found a heart")
	}
}

func GoStrRune() {
	runeCount("🫀a😌1🫀")
	runeCount("123")
	runeCount("a123")
}
