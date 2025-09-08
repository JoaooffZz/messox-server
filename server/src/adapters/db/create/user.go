package create

import (
	m "db/models"
	"strconv"
)

func (c *Create)User(user m.User)(string, error){
	var id int64
	query := `
	    INSERT INTO users
		(name, password, profile, bio)
		VALUES ($1, $2, $3, $4)
	`

	tx, err := c.DB.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	err = tx.QueryRow(query, user.Name, user.Password, user.Profile, user.Bio).Scan(&id)
	if err != nil {
		return "", err
	}
	
	err = tx.Commit()
	if err != nil {
		return "", nil
	}

	return strconv.Itoa(int(id)), nil
}