package userservice_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/shatwik7/polycrate/lib/db"
	userservice "github.com/shatwik7/polycrate/services/user_service"
	"github.com/stretchr/testify/assert"
)

var testDB *db.DB
var service *userservice.UserService

func setup() {
	// Setup a test DB connection
	var err error
	testDB, err = db.NewDB("postgres://polycrate:polycreate@localhost:5432/polycrate_db?sslmode=disable")
	if err != nil {
		panic(err)
	}
	service = userservice.NewUserService(testDB)
}

func teardown() {
	testDB.Exec("DELETE FROM user_credentials")
	testDB.Exec("DELETE FROM users")
	testDB.Close()
}

func TestCreateUser(t *testing.T) {
	setup()
	defer teardown()

	input := &userservice.CreateUserInput{
		Username:          "testuser",
		Email:             "test@example.com",
		FullName:          "Test User",
		ProfilePictureUrl: "",
		Bio:               "",
		Password:          "password123",
	}
	user, err := service.CreateUser(input)
	assert.NoError(t, err)
	assert.Equal(t, input.Username, user.Username)
	assert.Equal(t, input.Email, user.Email)
}

func TestGetUserByID(t *testing.T) {
	setup()
	defer teardown()

	input := &userservice.CreateUserInput{
		Username: "user2",
		Email:    "user2@example.com",
		Password: "pass2",
		FullName: "User Two",
		Bio:      "Bio",
	}
	createdUser, _ := service.CreateUser(input)
	found, err := service.GetUserByID(createdUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, createdUser.ID, found.ID)
}

func TestUpdateUser(t *testing.T) {
	setup()
	defer teardown()

	user, _ := service.CreateUser(&userservice.CreateUserInput{
		Username: "update_me",
		Email:    "update@site.com",
		Password: "pass",
	})

	update := &userservice.UpdateUserInput{
		ID:       user.ID,
		FullName: "Updated Name",
		Bio:      "Updated Bio",
	}
	updatedUser, err := service.UpdateUser(update)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedUser.FullName)
	assert.Equal(t, "Updated Bio", updatedUser.Bio)
}

func TestDeleteUser(t *testing.T) {
	setup()
	defer teardown()

	user, _ := service.CreateUser(&userservice.CreateUserInput{
		Username: "delete_me",
		Email:    "delete@site.com",
		Password: "pass",
	})
	ok, err := service.DeleteUser(user.ID)
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestLogin(t *testing.T) {
	setup()
	defer teardown()

	email := "login@site.com"
	password := "securepass"

	_, _ = service.CreateUser(&userservice.CreateUserInput{
		Username: "loginuser",
		Email:    email,
		Password: password,
	})

	user, err := service.Login(&userservice.LoginInput{
		Email:    email,
		Password: password,
	})
	assert.NoError(t, err)
	assert.Equal(t, email, user.Email)
}

func TestDeactivateUser(t *testing.T) {
	setup()
	defer teardown()

	user, _ := service.CreateUser(&userservice.CreateUserInput{
		Username: "inactive",
		Email:    "inactive@site.com",
		Password: "pass",
	})
	ok := service.DeactivateUser(user.ID)
	assert.True(t, ok)

	cred, _ := service.Repo.GetCredential(user.ID)
	assert.False(t, cred.IsActive)
}

func TestChangePassword(t *testing.T) {
	setup()
	defer teardown()

	user, _ := service.CreateUser(&userservice.CreateUserInput{
		Username: "changepass",
		Email:    "change@site.com",
		Password: "oldpass",
	})

	ok := service.ChangePassword(&userservice.ChangePasswordInput{
		ID:          user.ID,
		NewPassword: "newpass",
	})
	assert.True(t, ok)
}

func TestListUsers(t *testing.T) {
	setup()
	defer teardown()

	for i := 0; i < 5; i++ {
		_, _ = service.CreateUser(&userservice.CreateUserInput{
			Username: "user" + uuid.New().String(),
			Email:    uuid.New().String() + "@test.com",
			Password: "pass",
		})
	}
	list, err := service.ListUsers(10, 0)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(list), 5)
}
