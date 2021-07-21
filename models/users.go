package models

import (
	"errors"

	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

//variables block
var (
	//create pointers to the slice of users
	// will allow us to manipulate each users without
	// having to copy the user around
	users []*User
	//not specifying type will let Go compiler
	// guess the default type which is int
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not have an ID, or ID must be 0")
	}
	u.ID = nextID
	nextID++
	//we append to the slice of users a reference to the user's address in memory
	// since users holds pointers to the users
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with id '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
