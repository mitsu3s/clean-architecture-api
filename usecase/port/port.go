package port

import (
	"context"

	"github.com/mitsu3s/clean-architecture-api/entity"
)

// 図の赤色

type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
}

type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}

// userのCRUDに対するDB用のポート
type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}
