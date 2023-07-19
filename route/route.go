package route

import (
	"project/constant"
	"project/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	eJwt := e.Group("")
	eJwt.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(constant.SECRET_JWT),
	}))

	eJwt.GET("/users", controller.GetUsersController)
	eJwt.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)
	e.GET("/users/verification-email", controller.VerificationEmailUserController)
	eJwt.POST("/users/verification-email", controller.CreateVerificationEmailUserController)

	g := e.Group("/authentications")
	g.POST("/login", controller.LoginAuthenticationController)
	g.POST("/forgot-password", controller.ForgotPasswordController)
	g.POST("/change-password", controller.ChangePasswordController)

	return e
}
