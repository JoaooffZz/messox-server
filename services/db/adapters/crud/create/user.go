package create

import (
	"fmt"
	"os"
	ports "ports/db"

	"github.com/jackc/pgconn"
)

func (c *Create)NewUser(name string, password string)(*ports.User, error){

	if len(name) > 250 {
		return nil, &ports.StringLengthError{Field: name}
	}
	if len(password) > 250 {
		return nil, &ports.StringLengthError{Field: "password"}
	}

	path := os.Getenv("PROFILE_PATH")
	if path == "" {
		return nil, fmt.Errorf("profile path is empty")
	}
	profile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`
	    INSERT INTO %s
		(name, password, profile)
		VALUES ($1, $2, $3)
		RETURNING id
	`, tabelUsers)

	tx, err := c.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
    
	var id int
	err = tx.QueryRow(query, name, password, profile,).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" { // unique_violation
				return nil, &ports.ValidationError{
					Field: "Name: " + name,
					Msg:   "is already being used",
				}
			}
		}
		return nil, err
	}
	
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &ports.User{
		ID: id,
		Name: name,
		Password: password,
		Profile: profile,
		Bio: "",
	}, nil
}