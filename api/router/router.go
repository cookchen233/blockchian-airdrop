package router

import (
	"blockchain-deal-hunter/api/controller"
	"blockchain-deal-hunter/api/doc"
	"blockchain-deal-hunter/api/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programatically set swagger info
	doc.SwaggerInfo.Title = "swag_title_xxx"
	doc.SwaggerInfo.Description = "swag_descxxx"
	doc.SwaggerInfo.Version = "1.0"
	doc.SwaggerInfo.Host = "127.0.0.1:8880"
	doc.SwaggerInfo.BasePath = ""
	doc.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.New()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//非登陆接口
	store := sessions.NewCookieStore([]byte("secret"))
	r := router.Group("/api").Use(sessions.Sessions("mysession", store),
		middleware.RequestLogMiddleware(),
		middleware.RecoveryMiddleware(),
		middleware.TranslationMiddleware())
	{
		userController := controller.User{}
		r.POST("/login", userController.Login)
		r.GET("/login2", userController.Login2)
	}

	//登陆接口
	apiAuthGroup := router.Group("/api")
	apiAuthGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLogMiddleware(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
	}
	return router
}