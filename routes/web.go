package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	facades.Route().Static("/public", "./public")
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
			"hello":   "Goravel Framework",
		})
	})

	facades.Route().Get("/login", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("auth/login.tmpl", map[string]string{
			"title": "Login",
		})
	})
}
