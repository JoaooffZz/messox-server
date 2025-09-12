package create

import (
	m "db/models"
	"fmt"
)

func (c *Create)NewHistoryChat(chat m.Chat) error {
	
	query := fmt.Sprintf(`
	    INSERT INTO %s
		(chat_id, date, chat)
		VALUES ($1, $2, $3)
	`, tabelChats)

	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err  = tx.Exec(query, chat.ChatID, chat.Date, chat.Chat)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}