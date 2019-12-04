package models

// User is a user
type User struct {
	Name string
}

// NewUser creates a new User
func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
