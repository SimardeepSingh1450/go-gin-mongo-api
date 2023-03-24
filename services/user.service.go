package services

import "gin/mongo-api/models"

type UserService interface{
	//API Endpoint Functions/Controller Functions:
	CreateUser(*models.User) error
	GetUser(*string) (*models.User,error)
	GetAll() ([]*models.User,error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}