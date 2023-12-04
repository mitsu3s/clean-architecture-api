/*
- controller パッケージは，入力に対するアダプターです．

- ここでは，インプットポートとアウトプットポートを組み立てて，インプットポートを実行します．
- ユースケースレイヤからの戻り値を受け取って出力する必要はなく，純粋に http を受け取り，ユースケースを実行します．

- 図の緑色に該当.
*/

package controller

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/mitsu3s/clean-architecture-api/usecase/port"
)

type User struct {
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	// -> presenter.NewUserOutputPort
	InputFactory func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	// -> interactor.NewUserInputPort
	RepoFactory func(c *sql.DB) port.UserRepository
	// -> gateway.NewUserRepository
	Conn *sql.DB
}

// GetUserByIDは，httpを受け取り，全てのポート(UserInputPort, UserOutputPort, UserRepository)を組み立てて，inputPort.GetUserByID（usecaseのとこ）を呼び出す．
func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := strings.TrimPrefix(r.URL.Path, "/user/")
	outputPort := u.OutputFactory(w)
	repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetUserByID(ctx, userID)
}
