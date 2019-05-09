package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type CreateRequest struct {
	// API version
	Version  string `json:"v" query:"v"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Scope    string `form:"scope" binding:"required"`
}

type CreateResponse struct {
	// API version
	Version string       `json:"v"`
	Token   models.Token `json:"token"`
}

// TokenCreate godoc
// @Summary Create new token
// @Description Generate new access_token and refresh_token
// @ID create-token
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param email query string false "user email"
// @Param password query string false "user password"
// @Param scope query []string false "permissions, to check on authorization and request if necessary (Example: email,notif)"
// @Success 200 {object} handlers.CreateResponse
// @Failure 400 {object} handlers.ResponseData
// @Failure 500 {object} handlers.ResponseData
// @Router /token.create [Get]
func (h Handler) CreateHandler(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var user models.User
	if err := h.db.Where(&models.User{Email: r.Email}).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad password or email",
		})
		return
	}
	// сделал так как используется BCrypt на строне сервера
	if err := user.CheckPassword(r.Password); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad password or email",
		})
		return
	}

	jwttoken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := jwttoken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["role"] = user.Role
	claims["permissions"] = r.Scope

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := jwttoken.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusInternalServerError,
			Data:   "JWT signing error",
		})
		return
	}

	refresh, err := helpers.GenerateRandomString(128)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseData{
			Status: http.StatusInternalServerError,
			Data:   "Refresh token generate error",
		})
		return
	}

	token := models.Token{
		AccessToken:  t,
		RefreshToken: refresh,
		UserId:       1,
		Permissions:  strings.Split(r.Scope, ","),
	}

	h.db.Create(&token)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		Token:   token,
	})
}
