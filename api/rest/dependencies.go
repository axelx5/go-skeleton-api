package rest

import (
	"github.com/axelx5/go-skeleton-api/api/rest/handlers"
	"github.com/axelx5/go-skeleton-api/api/rest/handlers/account"
	"github.com/axelx5/go-skeleton-api/api/rest/handlers/ping"
	"github.com/axelx5/go-skeleton-api/application"
)

type HandlerRegistry struct {
	PingHandler          handlers.Handler
	CreateAccountHandler handlers.Handler
}

func wireHandlersDependencies() *HandlerRegistry {

	useCases := application.WireUseCaseDependencies()

	//Handlers
	handlers := HandlerRegistry{}

	handlers.PingHandler = &ping.Handler{}
	handlers.CreateAccountHandler = &account.Handler{
		CreateUseCase: useCases.CreateAccount,
	}

	return &handlers
}
