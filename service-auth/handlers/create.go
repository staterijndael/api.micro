package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/deissh/api.micro/service-auth/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type CreateRequest struct {
	// API version
	Api      string `json:"api" form:"api" query:"api"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type CreateResponse struct {
	// API version
	Api   string       `json:"api"`
	Token models.Token `json:"token"`
}

func (h Handler) CreateHandler(c echo.Context) error {
	// todo: verify user

	r := new(CreateRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
	}

	jwttoken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := jwttoken.Claims.(jwt.MapClaims)
	// todo: add params
	claims["name"] = "Jon Snow"
	//claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := jwttoken.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	refresh, err := helpers.GenerateRandomString(128)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseData{
			Status: http.StatusInternalServerError,
			Data:   "Refresh token generate error",
		})
	}

	token := models.Token{
		AccessToken:  t,
		RefreshToken: refresh,
		UserId:       1,
	}

	h.db.Create(&token)

	return c.JSON(http.StatusOK, CreateResponse{
		Api:   "v1",
		Token: token,
	})
}
