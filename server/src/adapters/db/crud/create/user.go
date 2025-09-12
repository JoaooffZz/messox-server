package create

import (
	m "db/models"
	"fmt"
)

func (c *Create)NewUser(user m.User)(int, error){
	query := fmt.Sprintf(`
	    INSERT INTO %s
		(name, password, profile, bio)
		VALUES ($1, $2, $3, $4)
	`, tabelUsers)

	tx, err := c.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
    
	var id int
	err = tx.QueryRow(query, user.Name, user.Password, user.Profile, user.Bio).Scan(&id)
	if err != nil {
		return 0, err
	}
	
	err = tx.Commit()
	if err != nil {
		return 0, nil
	}

	return id, nil
}