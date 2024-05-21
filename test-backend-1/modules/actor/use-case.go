package actor

import (
	"test-backend-1/entities"
	"test-backend-1/repositories"
	"test-backend-1/utils/hash"
)

type UseCaseInterface interface {
	validateActor(actor entities.Actor, validations ...validateFunc) (error, error)
	GetByID(id uint) (actor entities.Actor, err error)
	GetAllByUsername(username string, limit, offset uint) ([]entities.Actor, error)
	GetByUsername(username string) (actor entities.Actor, err error)
	CreateActor(actor *entities.Actor) (err error)
	ActivateActor(usernames string) error
	DeactivateActor(usernames string) error
	UpdateActor(actor *entities.Actor) (err error)
	DeleteActor(id uint) (err error)
}

func NewUseCase(
	repositoryInterface repositories.ActorRepositoryInterface,
	roleRepositoryInterface repositories.RoleRepositoryInterface,
) UseCaseInterface {
	return useCase{
		actorRepository: repositoryInterface,
		roleRepository:  roleRepositoryInterface,
	}
}

type useCase struct {
	actorRepository repositories.ActorRepositoryInterface
	roleRepository  repositories.RoleRepositoryInterface
}

func (uc useCase) validateActor(actor entities.Actor, validations ...validateFunc) (error, error) {
	for _, validation := range validations {
		validationError, err := validation(actor, uc.actorRepository)
		if err != nil {
			return nil, err
		}
		if validationError != nil {
			return validationError, nil
		}
	}
	return nil, nil
}

func (uc useCase) GetByID(id uint) (actor entities.Actor, err error) {
	actor, err = uc.actorRepository.GetByID(id)
	if err != nil {
		return
	}
	role, err := uc.roleRepository.GetByID(actor.RoleID)
	actor.Role = role
	return
}

func (uc useCase) GetByUsername(username string) (actor entities.Actor, err error) {
	actor, err = uc.actorRepository.GetByUsername(username)
	if err != nil {
		return
	}
	role, err := uc.roleRepository.GetByID(actor.RoleID)
	actor.Role = role
	return
}

func (uc useCase) GetAllByUsername(username string, limit, offset uint) ([]entities.Actor, error) {
	actors, err := uc.actorRepository.GetAllByUsername(username, limit, offset)
	return actors, err
}

func (uc useCase) CreateActor(actor *entities.Actor) (err error) {
	actor.RoleID = 1
	actor.Password, err = hash.Hash(actor.Password)
	if err != nil {
		return
	}
	if *actor, err = uc.actorRepository.Create(*actor); err != nil {
		return err
	}
	*actor, err = uc.GetByID(actor.ID)
	return
}

func (uc useCase) UpdateActor(actor *entities.Actor) (err error) {
	if actor.Password != "" {
		actor.Password, err = hash.Hash(actor.Password)
		if err != nil {
			return
		}
	}
	err = uc.actorRepository.Update(*actor)
	if err != nil {
		return
	}
	*actor, err = uc.GetByID(actor.ID)
	return
}

func (uc useCase) changeActiveActor(username string, value bool) error {
	actor, err := uc.actorRepository.GetByUsername(username)
	if err != nil {
		return err
	}
	actor.Active = value
	return uc.actorRepository.Save(actor)
}

func (uc useCase) ActivateActor(username string) error {
	return uc.changeActiveActor(username, true)
}

func (uc useCase) DeactivateActor(username string) error {
	return uc.changeActiveActor(username, false)
}

func (uc useCase) DeleteActor(id uint) (err error) {
	tx, err := uc.actorRepository.InitTransaction()
	if err != nil {
		return
	}
	actorTx := uc.actorRepository.Begin(tx)
	if err = actorTx.Delete(id); err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
