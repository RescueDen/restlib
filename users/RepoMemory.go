package users

import (
	"errors"
)

/**
Define a struct for Repo for use with users
*/
type RepoMemory struct {
	// The current id
	currentId int

	//A list of the sers
	usersList []User
}

//Provide a method to make a new UserRepoMemory
func NewRepoMemory() *RepoMemory {
	//Define a new repo
	newRepo := RepoMemory{
		0,
		make([]User, 0),
	}

	//Return a point
	return &newRepo

}

/**
Look up the user and return if they were found
*/
func (repo *RepoMemory) GetUserByEmail(email string) (User, error) {
	//March over each
	for _, v := range repo.usersList {
		//Check the email
		if v.Email == email {
			return v, nil
		}
	}

	return User{Id: -1}, errors.New("no user with email")
}

/**
Look up the user by id and return if they were found
*/
func (repo *RepoMemory) GetUser(id int) (User, error) {
	//March over each
	for _, v := range repo.usersList {
		//Check the email
		if v.Id == id {
			return v, nil
		}
	}

	return User{Id: -1}, errors.New("no user with id")
}

/**
Add the user to the database
*/
func (repo *RepoMemory) AddUser(t User) (User, error) {
	repo.currentId += 1
	t.Id = repo.currentId

	repo.usersList = append(repo.usersList, t)
	return t, nil
}

/**
Clean up the database, nothing much to do
*/
func (repo *RepoMemory) CleanUp() {

}

//func RepoDestroyCalc(id int) error {
//	for i, t := range usersList {
//		if t.Id == id {
//			usersList = append(usersList[:i], usersList[i+1:]...)
//			return nil
//		}
//	}
//	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
//}