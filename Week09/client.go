package main

import (
	"log"
	"net"
)

func main()  {
	log.Println("begin to dial...")
	conn, err := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	if err != nil {
		log.Println("Faild to dail ", err.Error())
		// handle error
		return
	}
	log.Println("Dail successfully")
}