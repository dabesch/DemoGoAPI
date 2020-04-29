package models

import (
	"errors"
	"fmt"
)

// User Structure
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // Using pointers so that it can be manipulated/shared throughout the application
	nextID = 1
)

//GetUsers retrurns the users struct
func GetUsers() []*User {
	return users
}

// AddUser adds the new User to the users struct
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil

}

// GetUserByID returns the User for the id provided
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

//UpdateUser the user with the id and replaces the current User in the users struct
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

//RemoveUserByID removes the user with the id from users
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
