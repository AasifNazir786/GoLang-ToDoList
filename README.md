# GoLang-ToDoList
# Todo List Application with PostgreSQL

This project is a simple **Todo List Application** that transitions from using an in-memory database to PostgreSQL for persistent data storage. It allows you to add, retrieve, update, and delete users, with all user information stored in a PostgreSQL database.

## Features

- **Add a User:** Allows adding new users to the system with details such as name, email, and address.
- **Retrieve All Users:** Fetches and displays a list of all users stored in the database.
- **Retrieve User by ID:** Retrieves the details of a specific user using their unique `user_id`.
- **Update a User:** Allows updating the user details based on `user_id`.
- **Delete a User:** Removes a user from the database by `user_id`.

## Prerequisites

- **Go 1.x**: The application is written in Go.
- **PostgreSQL**: The application uses PostgreSQL as the database to store user information.
- **PostgreSQL Driver (`lib/pq`)**: Go PostgreSQL driver is used to connect the app to the database.

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/AasifNazir786/GoLang-ToDoList.git
cd GoLang-ToDoList