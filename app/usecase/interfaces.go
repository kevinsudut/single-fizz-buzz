package usecase

import (
	"context"
)

type UsecaseItf interface {
	UseCaseSingleFizzBuzzWithRange(ctx context.Context, from int64, to int64) (resp string, err error)
}
