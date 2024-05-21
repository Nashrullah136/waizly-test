package actor

import (
	"github.com/gin-gonic/gin"
	"test-backend-1/middleware"
	"test-backend-1/repositories"
)

type Route struct {
	actorRequestHandler RequestHandlerInterface
}

func NewRoute(actorRepository repositories.ActorRepositoryInterface,
	roleRepository repositories.RoleRepositoryInterface,
) Route {
	useCase := NewUseCase(actorRepository, roleRepository)
	actorController := NewController(useCase)
	requestHandler := NewRequestHandler(actorController)
	return Route{actorRequestHandler: requestHandler}
}

func (r Route) Handle(router *gin.Engine) {
	router.PATCH("/me", middleware.Authenticate(), r.actorRequestHandler.UpdateProfile)
	actor := router.Group("/actors", middleware.Authenticate())
	actor.GET("/:id", r.actorRequestHandler.GetByID)
	actor.GET("", r.actorRequestHandler.GetAllByUsername)
	actor.PATCH("", middleware.AuthorizationWithRole([]string{"admin"}), r.actorRequestHandler.UpdateActor)
	actor.POST("", middleware.AuthorizationWithRole([]string{"admin"}), r.actorRequestHandler.CreateActor)
	actor.PATCH("/active", middleware.AuthorizationWithRole([]string{"admin"}), r.actorRequestHandler.ChangeActiveActor)
	actor.DELETE("/:id", middleware.AuthorizationWithRole([]string{"admin"}), r.actorRequestHandler.DeleteActor)
}
