package read

import (
	"database/sql"
	"fmt"
	m "ports/db"
)

func (r *Read) GetUser(name string) (*m.User, error) {
    query := fmt.Sprintf(`
        SELECT id, name, profile, password, bio
        FROM %s
        WHERE name = $1;
    `, tabelUsers)

    user := &m.User{}
    err := r.DB.QueryRow(query, name).Scan(&user.ID, &user.Name, &user.Profile, &user.Password, &user.Bio)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // usuário não encontrado
        }
        return nil, err // outro erro de DB
    }

    return user, nil
}
