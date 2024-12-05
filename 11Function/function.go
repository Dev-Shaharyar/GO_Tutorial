package function

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func fib(n int) int {
	var a, b int = 0, 1
	var c, sum int
	sum += a + b

	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
		sum += c
		//fmt.Println(c)
		fmt.Println(sum)
	}
	return sum
}

func addition(arr [5]int) int {

	sum := 0

	for _, val := range arr {
		sum += val
	}
	return sum
}

func Function() {
	var a, b int = 2, 5

	res := add(a, b)

	fmt.Println("2+5 =", res)

	sum := fib(10)

	fmt.Println("sum of the first 10 fib num is", sum)

	array := [5]int{1, 2, 3, 4, 5}

	ans := addition(array)

	fmt.Println(ans)

}
