package update

import "fmt"

func (u *Update)UpUserProfile(userID int, profile string) error {
	query := fmt.Sprintf(`
	    UPDATE %s
		SET
		    profile = $2
		WHERE id = $1;
	`, tabelUsers)

	_, err := u.DB.Exec(query, userID, profile)
    return err
}