package service

import (
	"Go-GitHub-Projects/GoLang-ToDoList/types"
	"fmt"
)

var Users []types.User

func GetUsers() []types.User {
	return Users
}

func GetUserById(id int) types.User {
	var user types.User
	for i := 0; i < len(Users); i++ {
		if Users[i].UserId == id {
			user = Users[i]
			break
		}
	}
	return user
}

func AddUser(user types.User) types.User {

	Users = append(Users, user)

	return user
}

func UpdateUser(id int, updatedUser types.User) types.User {
	var existingUser types.User
	var existingAddress types.Address
	for i := 0; i < len(Users); i++ {
		if Users[i].UserId == id {
			existingUser = Users[i]

			updatedAddress := updatedUser.UserAddress
			existingAddress.Country = updatedAddress.Country
			existingAddress.State = updatedAddress.State
			existingAddress.City = updatedAddress.City
			existingAddress.ZipCode = updatedAddress.ZipCode

			existingUser.UserId = id
			existingUser.FirstName = updatedUser.FirstName
			existingUser.LastName = updatedUser.LastName
			existingUser.UserEmail = updatedUser.UserEmail
			existingUser.UserAddress = existingAddress

			Users[i] = existingUser
			break
		}
	}
	return existingUser
}

func DeleteUser(id int) {
	Users = append(Users[:id], Users[id+1:]...)
	fmt.Printf("User with %d deleted successfully ", id)
}
