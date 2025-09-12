package read

import (
	m "db/models"
	"fmt"
)

func (r *Read)GetReceivedRequests(addresseeID int) ([]m.ViewUser, error){
	query := fmt.Sprintf(`
	    SELECT *,
		    user.name AS name, 
		    user.profile AS profile,
		    user.bio AS bio
	    FROM %s inbox
		JOIN users user ON inbox.sender_id = user.id
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