// No built in enum in GO

// IMplemented with the help of Type and Const
package enums

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota //0
	confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to ", status)
}
func Enums() {
	changeOrderStatus(Received)
}
