package main

type User struct {
	ID       int64
	Login    string
	Password string
	Salt     string
}

func (*User) TableName() string {
	return "users"
}
