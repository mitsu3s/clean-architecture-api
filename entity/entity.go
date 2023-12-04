/*
- entity パッケージは，ドメインモデルを実装します．．

- 図の黄色に該当.
*/

package entity

// ドメインロジックを実装する責務を持つ
// DB操作などの技術的な実装を持ってはならない
// また、他のどのレイヤにも依存してはならない

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
