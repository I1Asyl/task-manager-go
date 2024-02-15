package handler

import (
	"errors"

	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	services services.Service
}

func (a Auth) verifyUser(jwt string) bool {
	return a.services.VerifyUser(jwt)
}

func (a Auth) verifyAdmin(jwt string) bool {
	return a.services.VerifyAdmin(jwt)
}

func (a Auth) getUserIdByToken(token string) (int, error) {
	return a.services.GetUserIdByToken(token)
}

func NewAuth(services services.Service) *Auth {
	return &Auth{services: services}
}

func (a Auth) checkUser(ctx *gin.Context) {
	_, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("user id does not exist"))
	}
	ctx.JSON(200, gin.H{"message": "success"})
}
func (a Auth) login(ctx *gin.Context) {
	var user database.Model
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Given data is not corect"})
		return
	}
	access, refresh, mistakes, err := a.services.Login(user)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if len(mistakes) > 0 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Given data is not corect"})
		return
	}
	ctx.JSON(200, gin.H{"access": access, "refresh": refresh})
}
func (a Auth) refreshToken(ctx *gin.Context) {
	var user struct {
		Token string `json:"token"`
	}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Given data is not corect"})
		return
	}

	access, refresh, err := a.services.RefreshToken(user.Token)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, gin.H{"access": access, "refresh": refresh})
}

func (a Auth) logout(ctx *gin.Context) {
	var user struct {
		Token string `json:"token"`
	}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Given data is not corect"})
		return
	}
	if err := a.services.Logout(user.Token); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, gin.H{"message": "success"})
}

// func (a Auth) CreateUser(user database.User) error {
// 	return a.services.CreateUser(user)
// }
