package controllers

import (
	"fmt"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type LoginController struct {
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (r *LoginController) Index(ctx http.Context) http.Response {
	fmt.Println(facades.Config().GetString("http.port"))

	var users []models.User
	if err := facades.Orm().Query().Get(&users); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"code": 400,
			"msg":  err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"code":    200,
		"message": "success",
		"data":    users,
	})
}
