package services

import (
	"context"
	"errors"
	"github.com/nikolapetrovic1/ratemymusic/auth/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/auth/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/auth/pkg/utils"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/auth"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)
	user.Role = req.Role
	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}
	if user.Banned {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  "User banned",
		}, nil
	}
	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	accessToken, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status:      http.StatusOK,
		AccessToken: accessToken,
	}, nil
}

func (s *Server) Validate(_ context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.AccessToken)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	err = s.CheckRoles(req, claims.Role)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "user does not have needed role",
		}, nil
	}
	if user.Banned {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  "User banned",
		}, nil
	}
	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}

func (s *Server) CheckRoles(req *pb.ValidateRequest, userRole string) error {

	for _, role := range req.Roles {
		if role == userRole {
			return nil
		}
	}
	return errors.New("user does not have needed role")
}

func (s *Server) FindById(context context.Context, request *pb.IdRequest) (*pb.UserInfoResponse, error) {
	var user models.User
	s.H.DB.Where(&models.User{Id: request.Id}).First(&user)
	return &pb.UserInfoResponse{
		Id:    user.Id,
		Email: user.Email,
	}, nil
}

func (s *Server) GetAll(context context.Context, req *emptypb.Empty) (*pb.AllUsersResponse, error) {
	var users []models.User

	s.H.DB.Where("role = ?", "USER").Find(&users)
	return &pb.AllUsersResponse{
		Users: mapUserListToUserInfo(users),
	}, nil
}

func (s *Server) BanUser(context context.Context, request *pb.BanRequest) (*pb.AllUsersResponse, error) {

	s.H.DB.Model(&models.User{}).Where("id = ?", request.Id).Update("banned", request.Banned)
	return s.GetAll(context, &emptypb.Empty{})
}
func (s *Server) Delete(context context.Context, request *pb.IdRequest) (*pb.ValidateResponse, error) {
	s.H.DB.Delete(&models.User{}, request.Id)
	return &pb.ValidateResponse{
		Status: http.StatusAccepted,
		UserId: request.Id,
	}, nil
}

func (s *Server) Update(context context.Context, request *pb.EmailUpdateRequest) (*pb.UserInfoResponse, error) {
	s.H.DB.Model(&models.User{}).Where("email = ?", request.Email).Update("email", request.NewEmail)
	return &pb.UserInfoResponse{}, nil
}
func mapUserListToUserInfo(users []models.User) []*pb.UserInfoResponse {
	var userInfo []*pb.UserInfoResponse
	for _, user := range users {
		userInfo = append(userInfo, &pb.UserInfoResponse{
			Id:     user.Id,
			Email:  user.Email,
			Banned: user.Banned,
		})
	}
	return userInfo
}
