package account

import (
	"github.com/axelx5/go-skeleton-api/application/services/account"
)

type CreateUseCase struct {
	Service account.Service
}

func (uc CreateUseCase) Execute() error {
	uc.Service.Create()
	return nil
}
