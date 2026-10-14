package service

import (
	"context"
	
	"github.com/F4nk1/agentrix/backend/source/database"
	"github.com/F4nk1/agentrix/backend/source/model"
)
type Service interface {

}

type service struct {
	conn *database.Database
}

func NewService(conn *database.Database) Service {
	return &service{conn: conn}
}
