/*
- interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持ちます．
- interactorはアウトプットポートに依存し(importするということ)，インプットポートを実装します(interfaceを満たすようにmethodを追加するということ)．

- 図の赤色に該当．
*/

package interactor

import (
	"context"

	"github.com/mitsu3s/clean-architecture-api/usecase/port"
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   port.UserRepository
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}

// usecase.UserInputPortを実装している
// GetUserByID は，UserRepo.GetUserByIDを呼び出し，その結果をOutputPort.Render or OutputPort.RenderErrorに渡します．
func (u *User) GetUserByID(ctx context.Context, userID string) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.Render(user)
}
