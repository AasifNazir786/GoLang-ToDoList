package types

type User struct {
	UserId      int
	FirstName   string
	LastName    string
	UserEmail   string
	UserAddress Address
}

type Address struct {
	Country string
	State   string
	City    string
	ZipCode int
}
