package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Guilherme-AVVO/meu-primeiro-crud-go/src/configuration/rest_err"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomaininterface {
		return &userDomain{
			email, password, name, age,
		}
}

// regras de negocios
// domain nunca e exportado  
type userDomain struct {
	Email    string 
	Password string
	Name     string
	Age      int8   
}

func (ud *userDomain) EncryptPassword(){
	hash := md5.New()

	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

// Os metodos que o domain vai implementar
type UserDomaininterface interface{
	//recebe um obj do user e retorna se deu erro ou não
	CreateUser() *rest_err.RestErr
	//recebe uma string(id do user) e retorna se deu erro ou não
	UpdateUser(string) *rest_err.RestErr
	// recebe strig(id user) e retorna um ponteiro domain e err, caso de erro o domain precisa ser retorna como nil(por isso o ponteiro para retorno)
	FindUser(string) (*userDomain, *rest_err.RestErr)
	// recebe um obj do user e retorna se deu erro ou não
	DeleteUser(string) *rest_err.RestErr
}