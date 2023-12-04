/*
- port パッケージは，出力や入力などのポート(interface)を提供します．

- 図の赤色に該当．
*/

package port

import (
	"context"

	"github.com/mitsu3s/clean-architecture-api/entity"
)

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
