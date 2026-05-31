package model

import (
	"github.com/Guilherme-AVVO/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/Guilherme-AVVO/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)


func (ud *userDomain) CreateUser() *rest_err.RestErr{

	logger.Info("Init createUser model", zap.String("jouney","createUser"))
	
	ud.EncryptPassword()

	return nil
}