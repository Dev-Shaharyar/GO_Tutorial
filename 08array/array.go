package array

import "fmt"

func Array() {

	var a [5]int
	fmt.Println("empty", a) //[0 0 0 0 0]

	a[4] = 200
	fmt.Println("after seting the 4th index value", a) //[0 0 0 0 200]
	fmt.Println("get:", a[4])

	fmt.Println("length", len(a)) //5

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} //fill with zero which is left
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
