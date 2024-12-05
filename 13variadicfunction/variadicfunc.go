package variadicfunc

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	n := len(nums)
	// for i := 0; i < n; i++ {
	// 	total += nums[i]
	// }
	for i := range n {
		total += nums[i]
	}
	fmt.Println(total)
}

func VariadicFunc() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
