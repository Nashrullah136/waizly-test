package authentication

import (
	"fmt"
	"test-backend-1/dto"
	"test-backend-1/modules/actor"
	"test-backend-1/utils/hash"
	jwtUtil "test-backend-1/utils/jwt"
)

type ControllerInterface interface {
	Login(request LoginRequest) (dto.BaseResponse, error)
}

func NewAuthController(actorUseCase actor.UseCaseInterface) ControllerInterface {
	return controller{actorUseCase: actorUseCase}
}

type controller struct {
	actorUseCase actor.UseCaseInterface
}

func (c controller) Login(request LoginRequest) (dto.BaseResponse, error) {
	account, err := c.actorUseCase.GetByUsername(request.Username)
	defaultResponse := dto.ErrorUnauthorized("Wrong Username/Password")
	if err != nil {
		return defaultResponse, err
	}
	if err := hash.Compare(request.Password, account.Password); err != nil {
		return defaultResponse, err
	}
	if err := account.CanLogin(); err != nil {
		return dto.ErrorUnauthorized(err.Error()), nil
	}
	token, err := jwtUtil.GenerateJWT(account)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	result := LoginResponse{Token: fmt.Sprintf("Bearer %s", token)}
	return dto.Success("Authenticated success", result), nil
}
