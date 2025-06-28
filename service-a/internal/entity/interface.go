package entity

import "context"

type TempRepositoryInterface interface {
	Get(ctx context.Context, cep *Cep) (Temp, error)
}

