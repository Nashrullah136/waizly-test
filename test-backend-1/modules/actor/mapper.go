package actor

import (
	"test-backend-1/entities"
)

func mapCreateRequestToActor(request CreateRequest) entities.Actor {
	return entities.Actor{
		Username: request.Username,
		Password: request.Password,
		RoleID:   2,
	}
}

func mapUpdateRequestToActor(request UpdateRequest) entities.Actor {
	return entities.Actor{
		ID:       request.ID,
		Username: request.Username,
		Password: request.Password,
	}
}

func mapActorToResponse(actor entities.Actor) Representation {
	return Representation{
		Username: actor.Username,
		Role:     actor.Role.RoleName,
		Active:   actor.Active,
	}
}
