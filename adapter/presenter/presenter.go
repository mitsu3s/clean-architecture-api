/*
- presenter パッケージは，出力に対するアダプターです．

- ここでは，アウトプットポートを実装します(interfaceを満たすようにmethodを追加するということ)
*/

package presenter

import (
	"fmt"
	"net/http"

	"github.com/mitsu3s/clean-architecture-api/entity"
	"github.com/mitsu3s/clean-architecture-api/usecase/port"
)

type User struct {
	w http.ResponseWriter
}

// NewUserOutputPort はUserOutputPortを取得します．
func NewUserOutputPort(w http.ResponseWriter) port.UserOutputPort {
	return &User{
		w: w,
	}
}

// usecase.UserOutputPortを実装している
// Render はNameを出力します．
func (u *User) Render(user *entity.User) {
	u.w.WriteHeader(http.StatusOK)
	// httpでentity.User.Nameを出力
	fmt.Fprint(u.w, user.Name)
}

// RenderError はErrorを出力します．
func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(u.w, err)
}
