package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http/pprof"
	"payments-service/internal/repository"
	"payments-service/internal/server/http/handlers"
	"payments-service/internal/service"
)

type RoutePunishments struct {
	handler *handlers.PaymentsHandler
	router  *gin.RouterGroup
}

func NewPaymentsRoute(router *gin.RouterGroup, db *sqlx.DB) *RoutePunishments {
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	handler := handlers.NewPaymentsHandler(serv)

	return &RoutePunishments{
		handler: handler,
		router:  router,
	}
}

func (r *RoutePunishments) Routes(debug bool) {
	{
		r.router.POST("/create", r.handler.CreatePayment)
		r.router.GET("/accept", r.handler.AcceptPayment)
	}

	if debug {
		debugGroup := r.router.Group("/pprof")
		{
			debugGroup.GET("/index", gin.WrapF(pprof.Index))
			debugGroup.GET("/cmdline", gin.WrapF(pprof.Cmdline))
			debugGroup.GET("/profile", gin.WrapF(pprof.Profile))
			debugGroup.POST("/symbol", gin.WrapF(pprof.Symbol))
			debugGroup.GET("/symbol", gin.WrapF(pprof.Symbol))
			debugGroup.GET("/trace", gin.WrapF(pprof.Trace))
			debugGroup.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
			debugGroup.GET("/block", gin.WrapH(pprof.Handler("block")))
			debugGroup.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
			debugGroup.GET("/heap", gin.WrapH(pprof.Handler("heap")))
			debugGroup.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
			debugGroup.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
		}
	}
}
