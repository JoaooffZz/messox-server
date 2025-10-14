package read

import "fmt"


func (r *Read)GetTotalUsers() (int, error) {
	query := fmt.Sprintf(`
	    SELECT COUNT(*) 
		AS total_rows
		FROM %s;
	`, tabelUsers)

	var total int
	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}