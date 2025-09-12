package update

import "fmt"

func(u *Update)UpUserBio(userID int, bio string) error {
	query := fmt.Sprintf(`
	    UPDATE %s
		SET
		    bio = $2
		WHERE id = $1;
	`, tabelUsers)

	_, err := u.DB.Exec(query, userID, bio)
    return err
}