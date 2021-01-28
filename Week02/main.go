package main

import "fmt"

func main() {
	user, err := service.GetUserByID(1)
	if err != nil {
		fmt.Printf("original errorr is %T %v \n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack: %+v \n", err)
		return
	}
	fmt.Printf("User=%v\n", user)
}