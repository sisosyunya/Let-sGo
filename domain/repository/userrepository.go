package repository

import "FirstGo/graph/model"

// (引数)(返り値)
type UserRepository interface {
	CreateUser(input model.NewUser)(*model.User, error)
}