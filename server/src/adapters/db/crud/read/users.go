package read

import (
	m "db/models"
	"fmt"
)

func (r *Read)GetUsers(name string) ([]m.ViewUser, error) {
	query := fmt.Sprintf(`
	    SELECT 
		name, profile, bio
		FROM %s
		WHERE name LIKE $1;
	`, tabelUsers)

    rows, err := r.DB.Query(query, name+"%")

	if err != nil {
        return nil, err
    }
    defer rows.Close()

	var users []m.ViewUser

	for rows.Next() {
        var u m.ViewUser
        err := rows.Scan(&u.Name, &u.Profile, &u.Bio)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }

	return users, nil
}