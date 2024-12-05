package closure

import "fmt"

func x() func() {
	var a int = 10
	y := func() {
		fmt.Println(a)
	}
	return y
}

func Closure() {
	z := x()
	z()
}

//function with its lexical scope (Env) is called as Closure
