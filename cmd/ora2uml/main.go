package main

import (
	"fmt"

	"github.com/ehehalt/ora2uml"
)

func main() {
	fmt.Println("ora2uml starter")

	user := &ora2uml.ConfigUser{UserId: "mike", Password: "baum"}
	fmt.Println("UserId =", user.UserId)
}
