package userservice

import (
	"context"

	"github.com/google/uuid"
	"github.com/shatwik7/polycrate/lib/db"
	userpb "github.com/shatwik7/polycrate/lib/protos/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	Service UserService
}

func NewUserServer(database *db.DB) *UserServer {
	service := NewUserService(database)
	return &UserServer{Service: *service}
}

func convertUser(u User) *userpb.User {
	return &userpb.User{
		Id:                u.ID.String(),
		Username:          u.Username,
		Email:             u.Email,
		FullName:          u.FullName,
		ProfilePictureUrl: u.ProfilePictureUrl,
		Bio:               u.Bio,
		Website:           u.Website.String,
		Location:          u.Location.String,
		CreatedAt:         timestamppb.New(u.CreatedAt),
		UpdatedAt:         timestamppb.New(u.UpdatedAt),
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	input := &CreateUserInput{
		Username:          req.GetUsername(),
		Email:             req.GetEmail(),
		FullName:          req.GetFullName(),
		ProfilePictureUrl: req.GetProfilePictureUrl(),
		Bio:               req.GetBio(),
		Password:          req.GetPassword(),
	}
	user, err := s.Service.CreateUser(input)
	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{User: convertUser(*user)}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	user, err := s.Service.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserResponse{User: convertUser(*user)}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	input := &UpdateUserInput{
		ID:                id,
		FullName:          req.GetFullName(),
		ProfilePictureUrl: req.GetProfilePictureUrl(),
		Bio:               req.GetBio(),
	}
	user, err := s.Service.UpdateUser(input)
	if err != nil {
		return nil, err
	}
	return &userpb.UpdateUserResponse{User: convertUser(*user)}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	success, err := s.Service.DeleteUser(id)
	if err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{Success: success}, nil
}

func (s *UserServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := s.Service.ListUsers(int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		return nil, err
	}
	var pbUsers []*userpb.User
	for _, u := range users {
		pbUsers = append(pbUsers, convertUser(u))
	}
	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}

func (s *UserServer) SearchByEmail(ctx context.Context, req *userpb.SearchByEmailRequest) (*userpb.SearchByEmailResponse, error) {
	user, err := s.Service.SearchByEmail(req.GetEmail())
	if err != nil {
		return nil, err
	}
	return &userpb.SearchByEmailResponse{User: convertUser(*user)}, nil
}

func (s *UserServer) SearchByUsername(ctx context.Context, req *userpb.SearchByUsernameRequest) (*userpb.SearchByUsernameResponse, error) {
	users, err := s.Service.SearchByUserName(req.GetUsername(), int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		return nil, err
	}
	var pbUsers []*userpb.User
	for _, u := range users {
		pbUsers = append(pbUsers, convertUser(u))
	}
	return &userpb.SearchByUsernameResponse{Users: pbUsers}, nil
}

func (s *UserServer) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	input := &LoginInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	user, err := s.Service.Login(input)
	if err != nil {
		return nil, err
	}
	return &userpb.LoginResponse{
		User:  convertUser(*user),
		Token: "dummy-token", // replace with actual JWT if available
	}, nil
}

func (s *UserServer) Validate(ctx context.Context, req *userpb.ValidateRequest) (*userpb.ValidateResponse, error) {
	input := &LoginInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	valid := s.Service.Validate(input)
	return &userpb.ValidateResponse{Valid: valid}, nil
}

func (s *UserServer) ChangePassword(ctx context.Context, req *userpb.ChangePasswordRequest) (*userpb.ChangePasswordResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	input := &ChangePasswordInput{
		ID:          id,
		NewPassword: req.GetNewPassword(),
	}
	success := s.Service.ChangePassword(input)
	return &userpb.ChangePasswordResponse{Success: success}, nil
}

func (s *UserServer) DeactivateUser(ctx context.Context, req *userpb.DeactivateUserRequest) (*userpb.DeactivateUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	success := s.Service.DeactivateUser(id)
	return &userpb.DeactivateUserResponse{Success: success}, nil
}
