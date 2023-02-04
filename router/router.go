package router

import (
	"net/http"

	"MyGO.com/m/config"
	"MyGO.com/m/controller"
	"MyGO.com/m/repository"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDBConnection()

	//User
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func InitRoute() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	r.Use(Cors())

	r.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "Welcome to my API",
			"data":     nil,
		})
	})

	apiRoutes := r.Group("/api")

	//User routes
	userRoutes := apiRoutes.Group("user")
	{
		userRoutes.POST("/register", userController.Register)
	}

	panic(r.Run(":8090"))
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,Origin,X-Requested-With,Content-Type,Accept")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		c.Next()
	}
}