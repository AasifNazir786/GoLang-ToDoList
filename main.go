package main

import (
	"Go-GitHub-Projects/GoLang-ToDoList/service"
	"Go-GitHub-Projects/GoLang-ToDoList/types"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	// address := types.Address{
	// 	Country: "India",
	// 	State:   "Jammu And Kashmir",
	// 	City:    "Srinagar",
	// 	ZipCode: 190001,
	// }

	user1 := types.User{
		UserId:    1,
		FirstName: "Aasif Nazir",
		LastName:  "Shah",
		UserEmail: "aasif786@gmail.com",
		UserAddress: types.Address{
			Country: "India",
			State:   "Jammu And Kashmir",
			City:    "Srinagar",
			ZipCode: 190001,
		},
	}

	user2 := types.User{
		UserId:    2,
		FirstName: "Jamid Nazir",
		LastName:  "Shah",
		UserEmail: "jamid786@gmail.com",
		UserAddress: types.Address{
			Country: "India",
			State:   "Jammu And Kashmir",
			City:    "Srinagar",
			ZipCode: 190001,
		},
	}

	fmt.Println("Add User to in memory database.................")
	fmt.Println("===============================================================")
	printUser := service.AddUser(user1)
	fmt.Println(printUser)
	printUser1 := service.AddUser(user2)
	fmt.Println(printUser1)
	fmt.Println()
	fmt.Println("Retrieve all Users..............................")
	fmt.Println("===============================================================")
	users := service.GetUsers()
	for _, val := range users {
		fmt.Println(val)
	}
}
