package handler

import (
	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-gonic/gin"
)

type token struct {
	Token string `json:"token"`
}

type Auth struct {
	services services.Service
}

func NewAuth(services services.Service) *Auth {
	return &Auth{services: services}
}

// login godoc
// @Summary      Login
// @Description  Login user and return access and refresh tokens.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body database.Model  true  "User form"
// @Success      200  {object}  string
// @Failure      400  {object}  error
// @Router       /login [get]
func (a Auth) login(ctx *gin.Context) {
	var user database.Model
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	access, refresh, mistakes, err := a.services.Login(user)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if len(mistakes) > 0 {
		ctx.AbortWithStatusJSON(400, mistakes)
		return
	}
	ctx.JSON(200, gin.H{"access": access, "refresh": refresh})
}

// refreshToken godoc
// @Summary      Refresh token
// @Description  Recieves refresh token as a json named "token" and return access and refresh tokens.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        refresh body token true "Refresh token"
// @Success      200  {object}  string
// @Failure      400  {object}  error
// @Router       /refresh [get]
func (a Auth) refreshToken(ctx *gin.Context) {
	var user token
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

// func (a Auth) CreateUser(user database.User) error {
// 	return a.services.CreateUser(user)
// }
