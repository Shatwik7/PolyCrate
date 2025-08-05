package userservice

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shatwik7/polycrate/lib/db"
	"github.com/shatwik7/polycrate/services/user_service/auth"
)

type UserService struct {
	Repo *UserRepository
}

func NewUserService(database *db.DB) *UserService {
	UserRepo := NewUserRepository(database)
	return &UserService{Repo: UserRepo}
}

// ------------------- Create -------------------

func (service *UserService) CreateUser(u *CreateUserInput) (*User, error) {
	hashed, err := auth.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashed
	User, err := service.Repo.InsertUser(*u)
	if err != nil {
		return nil, err
	}
	UserCredential := &UserCredential{
		UserID:       User.ID,
		PasswordHash: hashed,
		LastLogin:    sql.NullTime{Time: time.Now(), Valid: true},
		IsActive:     true,
	}
	service.Repo.InsertCredential(*UserCredential)
	return User, nil
}

// ------------------- Get By ID -------------------

func (s *UserService) GetUserByID(id uuid.UUID) (*User, error) {
	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// ------------------- Update -------------------

func (s *UserService) UpdateUser(u *UpdateUserInput) (*User, error) {
	updatedUser, err := s.Repo.UpdateUser(*u)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// ------------------- Delete -------------------

func (s *UserService) DeleteUser(id uuid.UUID) (bool, error) {
	res, err := s.Repo.DeleteUser(id)
	return res, err
}

// ------------------- List All -------------------

func (s *UserService) ListUsers(limit int, offset int) ([]User, error) {
	Users, err := s.Repo.ListUsers(limit, offset)
	return Users, err
}

// ------------------- Find By Email -------------------

func (s *UserService) SearchByEmail(email string) (*User, error) {
	User, err := s.Repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return User, nil
}

// ------------------- FIND BY USERNAME -------------------

func (s *UserService) SearchByUserName(name string, limit int, offset int) ([]User, error) {
	return s.Repo.FindUsersByUsernamePartial(name, limit, offset)
}

// ------------------- ChangePassword -------------------

func (s *UserService) ChangePassword(ChangePasswordInput *ChangePasswordInput) bool {
	hashed, _ := auth.HashPassword(ChangePasswordInput.NewPassword)
	UpdateCred := &UserCredential{
		UserID:       ChangePasswordInput.ID,
		PasswordHash: hashed,
		LastLogin:    sql.NullTime{Time: time.Now(), Valid: true},
		IsActive:     true,
	}
	bol, err := s.Repo.UpdateCredential(*UpdateCred)
	if err != nil {
		return false
	}
	return bol
}

// ------------------- LOGIN -------------------

func (s *UserService) Login(u *LoginInput) (*User, error) {
	User, err := s.Repo.FindUserByEmail(u.Email)
	if err != nil {
		return nil, err
	}
	cred, err := s.Repo.GetCredential(User.ID)
	if err != nil {
		return nil, err
	}
	val := auth.CheckPasswordHash(u.Password, cred.PasswordHash)
	if !val {
		return nil, errors.New("unauthorized")
	}
	return User, nil
}

// ------------------- VALIDATE -------------------

func (s *UserService) Validate(u *LoginInput) bool {
	User, err := s.Repo.FindUserByEmail(u.Email)
	if err != nil {
		return false
	}
	cred, err := s.Repo.GetCredential(User.ID)
	if err != nil {
		return false
	}
	val := auth.CheckPasswordHash(u.Password, cred.PasswordHash)
	return val
}

// ------------------- DEACTIVATE -------------------

func (s *UserService) DeactivateUser(ID uuid.UUID) bool {
	UserCredential, err := s.Repo.GetCredential(ID)
	if err != nil {
		return false
	}
	UserCredential.IsActive = false
	res, err := s.Repo.UpdateCredential(*UserCredential)
	if err != nil {
		return false
	}
	return res
}
