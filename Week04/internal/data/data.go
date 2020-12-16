package data

type User struct {
	username string
	password string
}

type UserDAO interface{
	register(user User) error
}

func register(user User) error {
	// insert into db
	return nil
}