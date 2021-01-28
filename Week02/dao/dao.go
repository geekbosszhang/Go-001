package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	userID int
	userName string
}

func GetUserById(userID int) (user *User, err error){
	if id == 1 {
		return nil, errors.Wrap(sql.ErrNoRows, fmt.Sprintf("find user by id error, id %v", id))
	}
	return &User{
		userID: 1,
		userName: "hoho",
	}, nil
}

func main1() {
	res, err := query()
	if err != nil {
		fmt.Println("", errors.Cause)
	}
}

func query(sql string) ([]string, error) {
	result, err := excuteSQL(sql)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	if len(result) < 1 {
		return nil, err
	}
	return result, err
}

func excuteSQL(sql string) ([]string, error) {
	return nil, errors.New("no sql rows")
}