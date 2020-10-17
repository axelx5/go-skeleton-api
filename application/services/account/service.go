package account

import (
	"log"
)

type Service struct {
}

func (Service) Create() error {
	log.Println(`Use case create account executed`)
	return nil
}
