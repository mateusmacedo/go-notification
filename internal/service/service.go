package service

import (
	"github.com/mateusmacedo/go-notification/internal/repository"
	"github.com/mateusmacedo/go-notification/pkg/helper/sid"
	"github.com/mateusmacedo/go-notification/pkg/jwt"
	"github.com/mateusmacedo/go-notification/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(tm repository.Transaction, logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
