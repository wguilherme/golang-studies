package main

import (
	"log/slog"
)

type User struct {
	ID       int64 `json:"id,string"`
	Username string
	Role     string
	Password Password `json:"password"`
}

type Password string

// LogValue implementa a interface LogValuer do slog
// Essa interface permite customizar como o tipo será logado
// Em vez de usar a representação padrão da struct, podemos:
// 1. Escolher quais campos mostrar
// 2. Modificar valores sensíveis (como senhas)
// 3. Formatar os dados de maneira específica
// 4. Adicionar campos calculados
func (u User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int64("id", u.ID),
		slog.String("username", u.Username),
		slog.String("role", u.Role),
		slog.String("password", "[REDACTED]"),
	)
}

func main() {

	u := User{
		Username: "admin",
		ID:       1,
		Role:     "admin",
		Password: "123456",
	}

	// Ao chamar slog.Info com nossa struct User,
	// o slog automaticamente usa o método LogValue
	// Em vez de logar todos os campos da struct diretamente,
	// ele usa nossa implementação customizada que:
	// - Mantém ID, Username e Role como estão
	// - Substitui a senha real por [REDACTED]
	//
	// Resultado no log:
	// INFO user u.id=1 u.username=admin u.role=admin u.password=[REDACTED]

	slog.Info("user", "u", u)

}
