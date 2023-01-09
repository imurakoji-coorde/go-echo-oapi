package main

import (
	"sample/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ServerInterface の定義を満たす構造体
type server struct{}

// ListUsers は ServerInterface に定義される ListUsers の実装
func (h server) ListUsers(ctx echo.Context, params api.ListUsersParams) error {
	// os.Stdout
	return ctx.JSON(200, map[string]interface{}{
		"msg":   "200 success",
		"limit": params.Limit,
		"page":  params.Page,
	})
}

// CreateUsers は ServerInterface　に定義される CreateUsers の実装
func (h server) CreateUsers(ctx echo.Context) error {
	b := &api.CreateUsersJSONBody{}
	// b.Name = ctx.FormValue("name")
	if err := ctx.Bind(b); err != nil {
		return err
	}

	return ctx.JSON(
		200,
		map[string]interface{}{
			"msg":     "200 success",
			"reqBody": b,
		})

}

// ServerInterface の定義を満たす server をルーティングする
// https://github.com/deepmap/oapi-codegen#registering-handlers の Echo の箇所を参考にします
func setHandler() *echo.Echo {
	s := server{}
	e := echo.New()
	e.Use(middleware.CORS())

	// RegisterHandlers の第2引数の型は ServerInterface なので、
	// ServerInterface を満たした構造体を代入する
	api.RegisterHandlers(e, s)

	return e
}

func main() {
	e := setHandler()
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
