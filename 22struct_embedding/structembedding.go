package structembedding

import "fmt"

type order struct {
	id      string
	ammount float32
	status  string
	//createdAt time.Time //nonosecond precision
	customer //struct embedding
}

type customer struct {
	name  string
	phone string
}

func Structembedding() {

	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "9458200368",
	// }
	newOrder := order{
		id:      "1",
		ammount: 30,
		status:  "received",
		//customer: newCustomer, //adding the data of the customer
		customer: customer{
			name:  "john",
			phone: "935832657",
		},
	}
	newOrder.customer.name = "robin" //changing the name
	fmt.Println(newOrder)
}
