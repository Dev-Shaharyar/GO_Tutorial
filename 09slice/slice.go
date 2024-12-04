package slice

import (
	"fmt"
	"slices"
	"sort"
)

func Slice() {

	var s = []string{}
	fmt.Println("type of the s is", s)

	fmt.Println("uninit:", s, s == nil, len(s) == 0) // [] , true , true

	s = make([]string, 3)

	fmt.Println("empty:", s, "length", len(s), "capacity", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("values of s is ", s)
	fmt.Println("value at s[2] is ", s[2])

	fmt.Println("length", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("after apended value of s is ", s) //[a,b,c,d,e,f]

	c := make([]string, len(s)) //[type , length]
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]            //[include : exclude]
	fmt.Println("sl1:", l) //[c,d,e]

	l = s[:5]              //start from 0 indx
	fmt.Println("sl2:", l) //[a,b,c,d,e]

	l = s[2:]              // till last
	fmt.Println("sl3:", l) //c,d,e,f

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	score := make([]int, 4)

	score[0] = 200

	score[1] = 150

	score[2] = 100

	score[3] = 175

	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score)) //false

	sort.Ints(score) //sorting in incr order

	fmt.Println(score)

	fmt.Println(sort.IntsAreSorted(score)) //true

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
