package magazine

// Subscriber Structure
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

// Employee Structure
type Employee struct {
	Name   string
	Salary float64
	Address
}

// Address Structure
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
