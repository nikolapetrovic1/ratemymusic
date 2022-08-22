package services

import (
	"context"
	"errors"
	"github.com/nikolapetrovic1/ratemymusic/auth/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/auth/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/auth/pkg/utils"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/auth"
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
