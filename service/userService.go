package service

import (
	"Go-GitHub-Projects/GoLang-ToDoList/types"
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

// InitDB initializes the connection to the PostgreSQL database.
func InitDB() {
	connStr := "name=postgres password=njasm786 dbname=todo-db sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Database connection is not alive: %v", err)
	}
	fmt.Println("Connected to the database!")
}

// this method adds new users into databases
func AddUser(user types.User) error {
	query := `INSERT INTO users (first_name, last_name, email, country, state, city, zipcode)
				VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id`

	var id int

	err := db.QueryRow(query, user.FirstName, user.LastName, user.UserEmail, user.UserAddress.Country,
		user.UserAddress.State, user.UserAddress.City, user.UserAddress.ZipCode).Scan(&id)

	if err != nil {
		return err
	}

	user.UserId = id
	return nil
}

// this method retrieves all the users present in databases
func GetAllUsers() ([]types.User, error) {
	query := `SELECT * FROM users`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.User
	for rows.Next() {

		var user types.User
		var address types.Address

		if err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.UserEmail,
			&address.Country, &address.State, &address.City, &address.ZipCode); err != nil {
			return nil, err
		}

		user.UserAddress = address
		users = append(users, user)
	}
	return users, nil
}

// retrieve a specific user by its ID
func GetUserById(id int) (types.User, error) {
	query := `SELECT * FROM users WHERE user_id = $1`

	var user types.User
	var address types.Address

	err := db.QueryRow(query, id).Scan(
		&user.UserId, &user.FirstName, &user.LastName, &user.UserEmail,
		&address.Country, &address.State, &address.City, &address.ZipCode)
	if err != nil {
		return types.User{}, err
	}
	user.UserAddress = address
	return user, nil
}

// this method updates the existing user
func UpdateUser(id int, updatedUser types.User) error {
	query := `
		UPDATE users
		SET first_name = $1, last_name = $2, email = $3, country = $4, state = $5, city = $6, zipcode = $7
		WHERE user_id = $8
		`
	_, err := db.Exec(query, query, updatedUser.FirstName, updatedUser.LastName, updatedUser.UserEmail, updatedUser.UserAddress.Country,
		updatedUser.UserAddress.State, updatedUser.UserAddress.City, updatedUser.UserAddress.ZipCode, id)

	return err
}

// this method deletes the user from database
func DeleteUser(userId int) error {
	query := `DELETE FROM users WHERE user_id = $1;`
	_, err := db.Exec(query, userId)
	return err
}

func CloseDB() {
	if db != nil {
		_ = db.Close()
		fmt.Println("Database connection closed.")
	}
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
