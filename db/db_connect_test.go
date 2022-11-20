package db

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestAddUser(t *testing.T) {
	// begin by creating the mock database we will use for the tests
	DbMock(CREATE)
	// wait until we are finished in this function then drop the mock-database
	defer DbMock(DROP)

	//initialize a new user and use AddDBUser to add it into our database
	user := "mvc_test"
	AddDBUser(user)

	db, err := DbConnect()

	if err != nil {
		t.Errorf("Error connecting to database in %v\n%v", t.Name(), err)
	}
	defer db.Close()

	var newUserInDB DBUser

	userNotFound := db.Where("username = ?", user).First(&newUserInDB).RecordNotFound()

	if userNotFound {
		t.Errorf("%v was addded but was not retrieved.", user)
	}

}
