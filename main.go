package main

import (
	"Go-GitHub-Projects/GoLang-ToDoList/service"
	"Go-GitHub-Projects/GoLang-ToDoList/types"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	//create an Object of Address struct
	address := types.Address{
		Country: "India",
		State:   "Jammu And Kashmir",
		City:    "Srinagar",
		ZipCode: 190001,
	}

	//create an object of User struct and also embed Address into it
	user1 := types.User{
		UserId:      1,
		FirstName:   "Aasif Nazir",
		LastName:    "Shah",
		UserEmail:   "aasif786@gmail.com",
		UserAddress: address,
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

	fmt.Println("Add User to in-memory database.................")
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
	fmt.Println()

	fmt.Println("Retrieve data from database using id...................")
	fmt.Println("===============================================================")
	id := 1
	fmt.Println(service.GetUserById(id))
	fmt.Println()

	fmt.Println("Update User at a specific id.....................")
	fmt.Println("===============================================================")
	updateUser := types.User{
		UserId:    101,
		FirstName: "Aabid Hamid",
		LastName:  "Shah",
		UserEmail: "aabid786@gmail.com",
		UserAddress: types.Address{
			Country: "India",
			State:   "Jammu And Kashmir",
			City:    "Srinagar",
			ZipCode: 192301,
		},
	}
	fmt.Println(service.UpdateUser(1, updateUser))
	fmt.Println()

	fmt.Println("Retrieve all Users after updation..............................")
	fmt.Println("===============================================================")
	users1 := service.GetUsers()
	for _, val := range users1 {
		fmt.Println(val)
	}
	fmt.Println()

	// Delete user from in-memory database
	service.DeleteUser(1)

	fmt.Println("Retrieve all Users after deletion..............................")
	fmt.Println("===============================================================")
	users2 := service.GetUsers()
	for _, val := range users2 {
		fmt.Println(val)
	}
	fmt.Println()

}
