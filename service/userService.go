package service

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func init() {
	connStr := "name=postgres password=njasm786 dbname=todo-db sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection is not alive: %v", err)
	}
	fmt.Println("Connected to the database!")
}

// // In-memory database of users
// var Users []types.User

// // GetUsers retrieves all users from the in-memory database.
// func GetUsers() []types.User {
// 	return Users
// }

// // GetUserById retrieves a user by their ID from the in-memory database.
// func GetUserById(id int) types.User {
// 	var user types.User
// 	for i := 0; i < len(Users); i++ {
// 		if Users[i].UserId == id {
// 			user = Users[i]
// 			break
// 		}
// 	}
// 	return user
// }

// // AddUser adds a new user to the in-memory database and returns the added user.
// func AddUser(user types.User) types.User {
// 	Users = append(Users, user)
// 	return user
// }

// // Update a User to the in-memory database and returns the updated user.
// func UpdateUser(id int, updatedUser types.User) types.User {
// 	var existingUser types.User
// 	var existingAddress types.Address
// 	for i := 0; i < len(Users); i++ {
// 		if Users[i].UserId == id {
// 			existingUser = Users[i]

// 			updatedAddress := updatedUser.UserAddress
// 			existingAddress.Country = updatedAddress.Country
// 			existingAddress.State = updatedAddress.State
// 			existingAddress.City = updatedAddress.City
// 			existingAddress.ZipCode = updatedAddress.ZipCode

// 			existingUser.UserId = id
// 			existingUser.FirstName = updatedUser.FirstName
// 			existingUser.LastName = updatedUser.LastName
// 			existingUser.UserEmail = updatedUser.UserEmail
// 			existingUser.UserAddress = existingAddress

// 			Users[i] = existingUser
// 			break
// 		}
// 	}
// 	return existingUser
// }

// // Delete user from an in-memory database
// func DeleteUser(id int) {
// 	if id < len(Users)-1 && id >= 0 {
// 		Users = append(Users[:id], Users[id+1:]...)
// 		fmt.Printf("User with %d deleted successfully ", id)
// 	} else if id == len(Users) {
// 		Users = append(Users[0 : id-1])
// 		fmt.Printf("User with %d deleted successfully ", id)
// 	} else {
// 		fmt.Println("Can't delete user because id is not present in database")
// 	}
// }
