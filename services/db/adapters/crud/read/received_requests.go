package read

import (
	"fmt"
	m "ports/db"
)

func (r *Read)GetReceivedRequests(addresseeID int) ([]m.ViewUser, error){
	query := fmt.Sprintf(`
		SELECT
			u.name AS name, 
			u.profile AS profile,
			u.bio AS bio
		FROM %s inbox
		JOIN users u ON inbox.sender_id = u.id
		WHERE inbox.addressee_id = $1;
	`, tabelInboxRequests)

	rows, err := r.DB.Query(query, addresseeID)

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