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

	// user rest api

	eJwt.GET("/users", controller.GetUsersController)
	eJwt.PUT("/users/:id/password", controller.UpdatePasswordUserController)
	eJwt.GET("/users/filter", controller.GetFilterUsersController)
	eJwt.GET("/users/gender/:gender", controller.GetByGenderUsersController)
	eJwt.GET("/users/method/:method", controller.GetByRegistrationMethodUserController)
	eJwt.GET("/users/verified-email/:status", controller.GetByVerifiedEmailStatusUsersController)
	eJwt.GET("/users/birth-year/:year", controller.GetByBirthYearUsersController)
	eJwt.GET("/users/empty-profile-photo", controller.GetEmptyProfilePhotoUsersController)
	eJwt.GET("/users/token-expired", controller.GetTokenExpiredUsersController)
	eJwt.GET("/users/token-verified-email/:token", controller.GetTokenVerifiedEmailUsersController)
	eJwt.GET("/users/joined-date-range", controller.GetJoinedDateRangeUsersController)
	eJwt.GET("/users/search", controller.GetSearchUsersController)
	eJwt.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)
	e.GET("/users/verification-email", controller.VerificationEmailUserController)
	eJwt.POST("/users/verification-email", controller.CreateVerificationEmailUserController)

	e.POST("/questions", controller.CreateQuestionController)
	e.POST("/questions/multiple-choice", controller.CreateQuestionByMultipleChoiceController)
	e.POST("/questions/true-false", controller.CreateQuestionByTrueFalseController)
	e.POST("/questions/fill-in", controller.CreateQuestionByFillInController)
	eJwt.GET("/questions", controller.GetQuestionsController)
	eJwt.GET("/questions/all", controller.GetByTypeQuestionsController)
	eJwt.GET("/questions/:user_id/:quiz_id", controller.GetByUserIDQuizIDQuestionController)

	// question category rest api
	eJwt.GET("/questions-categories", controller.GetQuestionCategoriesController)
	eJwt.GET("/questions-categories/:id", controller.GetQuestionCategoryController)
	eJwt.POST("/questions-categories", controller.CreateQuestionCategoryController)
	eJwt.PUT("/questions-categories/:id", controller.UpdateQuestionCategoryController)
	eJwt.DELETE("/questions-categories/:id", controller.DeleteQuestionCategoryController)

	// wallet rest api
	eJwt.GET("/wallets", controller.GetWalletsController)
	eJwt.GET("/wallets/:id", controller.GetWalletController)
	eJwt.POST("/wallets", controller.CreateWalletController)
	eJwt.PUT("/wallets/:id", controller.UpdateWalletController)
	eJwt.DELETE("/wallets/:id", controller.DeleteWalletController)
	eJwt.POST("/wallets/send", controller.SendWalletController)

	// quiz rest api
	eJwt.GET("/quizzes", controller.GetQuizzesController)
	eJwt.GET("/quizzes/:id", controller.GetQuizController)
	eJwt.GET("/quizzes/:user_id", controller.GetByUserIDQuizController)
	eJwt.GET("/quizzes/:user_id/:quiz_id", controller.GetByUserIDUserQuizQuizController)
	eJwt.POST("/quizzes", controller.CreateQuizController)
	eJwt.PUT("/quizzes/:id", controller.UpdateQuizController)
	eJwt.DELETE("/quizzes/:id", controller.DeleteQuizController)

	// level rest api
	eJwt.GET("levels", controller.GetLevelsController)
	eJwt.GET("levels/:id", controller.GetLevelController)
	eJwt.POST("levels", controller.CreateLevelController)
	eJwt.PUT("levels/:id", controller.UpdateLevelController)
	eJwt.DELETE("levels/:id", controller.DeleteLevelController)

	// package rest api
	eJwt.GET("/packages", controller.GetPackagesController)
	eJwt.GET("/packages/:id", controller.GetPackageController)
	eJwt.POST("/packages", controller.CreatePackageController)
	eJwt.PUT("/packages/:id", controller.UpdatePackageController)
	eJwt.DELETE("/packages/:id", controller.DeletePackageController)

	g := e.Group("/authentications")
	g.POST("/login", controller.LoginAuthenticationController)
	g.POST("/forgot-password", controller.ForgotPasswordController)
	g.POST("/change-password", controller.ChangePasswordController)

	return e
}
