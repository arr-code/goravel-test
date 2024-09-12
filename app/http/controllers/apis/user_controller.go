package apis

import (
	"fmt"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (r *UserController) Index(ctx http.Context) http.Response {
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

func (r *UserController) Show(ctx http.Context) http.Response {
	var user models.User
	if err := facades.Orm().Query().Where("id", ctx.Request().Input("id")).First(&user); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"code":    400,
			"message": err.Error(),
		})
	}
	return ctx.Response().Success().Json(http.Json{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}

func (r *UserController) Store(ctx http.Context) http.Response {
	user := models.User{
		Name:     ctx.Request().Input("name"),
		Email:    ctx.Request().Input("email"),
		Password: ctx.Request().Input("password"),
	}

	if err := facades.Orm().Query().Create(&user); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"code":    400,
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}

func (r *UserController) Update(ctx http.Context) http.Response {
	var user models.User
	if err := facades.Orm().Query().Where("id", ctx.Request().Input("id")).FirstOrFail(&user); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"code":    400,
			"message": err.Error(),
		})
	}

	user.Name = ctx.Request().Input("name")
	user.Email = ctx.Request().Input("email")

	if err := facades.Orm().Query().Save(&user); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"code":    400,
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}

func (r *UserController) Destroy(ctx http.Context) http.Response {
	var user models.User
	result, err := facades.Orm().Query().Where("id", ctx.Request().Input("id")).Delete(&user)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"code":    400,
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"code":    200,
		"message": "success",
		"data":    result,
	})
}
