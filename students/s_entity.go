package students

type Students struct {
	ID      int
	Name    string
	Job     string
	Reason  string
	Address Address
}

type Address struct {
	Street string
	City   string
}
