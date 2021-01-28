package service

import "dao"

func GetUserByID(id int) (*dao.User, error) {
	user, err := dao.GetUserByID(id)
	if err == nil {
		return user, nil
	}
	return nil, errors.Wrap(err, "UserService GetUserID failed")
}