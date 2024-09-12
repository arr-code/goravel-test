package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers/apis"
)

func Api() {
	userController := apis.NewUserController()
	facades.Route().Resource("/api/users", userController)
}
