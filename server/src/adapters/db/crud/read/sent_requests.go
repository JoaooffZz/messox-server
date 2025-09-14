package read

import (
	m "db/models"
	"fmt"
)

func(r *Read)GetSentRequests(senderID int) ([]m.ViewUser, error) {
	query := fmt.Sprintf(`
	    SELECT
		    u.name AS name, 
		    u.profile AS profile,
		    u.bio AS bio
	    FROM %s inbox
		JOIN users u ON inbox.addressee_id = u.id
		WHERE inbox.sender_id = $1;
	`, tabelInboxRequests)

	rows, err := r.DB.Query(query, senderID)

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