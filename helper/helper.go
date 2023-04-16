package helper

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) string {
	salt := 8
	password := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func ComparePassword(loginPassword, localPassword string) error {
	login, local := []byte(loginPassword), []byte(localPassword)
	err := bcrypt.CompareHashAndPassword(local, login)
	return err
}

func GetContentType(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}
