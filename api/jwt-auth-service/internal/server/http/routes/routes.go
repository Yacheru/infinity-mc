package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-auth-service/internal/jwt"
	"jwt-auth-service/internal/server/http/middleware"

	_ "jwt-auth-service/docs"
	"jwt-auth-service/init/config"
	"jwt-auth-service/internal/repository"
	"jwt-auth-service/internal/server/http/handlers"
	"jwt-auth-service/internal/service"
	"jwt-auth-service/pkg/email"
	"jwt-auth-service/pkg/utils"
)

type Router struct {
	router  *gin.RouterGroup
	handler *handlers.Handlers

	jwtManager *jwt.Manager
}

func InitRouterAndComponents(router *gin.RouterGroup, pdb *sqlx.DB, r *redis.Client, coll *mongo.Collection, cfg *config.Config) *Router {
	repo := repository.NewRepository(pdb, r, coll, cfg)
	jwtManager := jwt.NewJWTManager(cfg.Salt)
	hasher := utils.NewSHA512Hasher(cfg.Salt)
	mail := email.NewSMPTServer(cfg.EmailSender, cfg.EmailSenderPassword, cfg.SmtpHost, cfg.SmtpPort)
	serv := service.NewService(repo, jwtManager, mail, cfg, hasher)
	handler := handlers.NewHandlers(serv, jwtManager)

	return &Router{
		router:     router,
		handler:    handler,
		jwtManager: jwtManager,
	}
}

func (r *Router) Routes() {
	{
		r.router.POST("/send-code", r.handler.SendCode)
		r.router.POST("/register", middleware.ValidateVerifCode(), r.handler.Register)
		r.router.POST("/login", r.handler.Login)
		r.router.POST("/refresh", r.handler.RefreshTokens)
		r.router.GET("/logout", r.handler.Logout)
	}

	auth := r.router.Group("/", middleware.Auth(r.jwtManager), middleware.IsAdmin())
	{
		auth.GET("/users", r.handler.GetAll)
		auth.PATCH("/update-role/:uuid", middleware.ValidateUUID(), middleware.ValidateRole(), r.handler.UpdateRole)
		auth.DELETE("/delete-user/:uuid", middleware.ValidateUUID(), r.handler.DeleteUser)

	}

	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
