package application

import (
	_accountServices "github.com/axelx5/go-skeleton-api/application/services/account"
	"github.com/axelx5/go-skeleton-api/application/usecases"
	_accountUseCases "github.com/axelx5/go-skeleton-api/application/usecases/account"
)

type UseCaseRegistry struct {
	CreateAccount usecases.UseCase
}

func WireUseCaseDependencies() *UseCaseRegistry {

	useCaseRegistry := UseCaseRegistry{}

	accountService := _accountServices.Service{}
	useCaseRegistry.CreateAccount = &_accountUseCases.CreateUseCase{
		Service: accountService,
	}

	return &useCaseRegistry
}
