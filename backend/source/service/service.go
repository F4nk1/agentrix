package service

import (
	
	"github.com/F4nk1/agentrix/backend/source/database"
)
type Service interface {

}

type service struct {
	conn *database.Database
}

func NewService(conn *database.Database) Service {
	return &service{conn: conn}
}
