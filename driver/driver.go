/*
- driver パッケージは，技術的な実装を持ちます．
*/

package driver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitsu3s/clean-architecture-api/adapter/controller"
	"github.com/mitsu3s/clean-architecture-api/adapter/gateway"
	"github.com/mitsu3s/clean-architecture-api/adapter/presenter"
	"github.com/mitsu3s/clean-architecture-api/usecase/interactor"
)

// driver.goではDBのconnectionを生成し，routingの設定．
// controller.goで定義されているcontroller.Userを作成し，http.HandleFunc()にcontroller.User.GetUserByIDを渡している.

// Serve はserverを起動させる．
func Serve(addr string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DATABASE"))
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	// Controller（controllerのとこ）の定義から作成
	user := controller.User{
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		Conn:          conn,
	}

	http.HandleFunc("/user/", user.GetUserByID)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
