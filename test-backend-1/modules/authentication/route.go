package authentication

import (
	"github.com/gin-gonic/gin"
	"test-backend-1/modules/actor"
)

type Route struct {
	authRequestHandler RequestHandlerInterface
}

func NewRoute(actorUseCase actor.UseCaseInterface) Route {
	controller := NewAuthController(actorUseCase)
	requestHandler := NewRequestHandler(controller)
	return Route{
		authRequestHandler: requestHandler,
	}
}

func (r Route) Handle(router *gin.Engine) {
	router.POST("/login", r.authRequestHandler.Login)
}
