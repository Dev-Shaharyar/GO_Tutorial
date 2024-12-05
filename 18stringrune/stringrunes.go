package stringrunes

import (
	"fmt"
	"unicode/utf8"
)

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

func Stringrunes() {
	//internally in go string are just collection of byte
	//and bytes are just is alias of unit8

	var s = "Iamआमt"

	fmt.Println("leng:", len(s))

	fmt.Println([]byte(s))

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("value = %v, Type = %T\n", s[i], s[i]) //print byte in hexa
	// }

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%c starts at %d\n", runeValue, idx)
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%c starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}
