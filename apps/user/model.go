package main

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
}

func (*User) TableName() string {
	return "users"
}

/*
docker run --rm -d -p 5432:5432 --name postgres --env POSTGRES_DB=mydb --env POSTGRES_USER=myuser --env POSTGRES_PASSWORD=mypassword postgres:latest

psql -h localhost -p 5432 -U myuser -W mydb

create table users (id bigserial primary key, first_name varchar, last_name varchar, login varchar, password varchar, salt varchar);
*/
