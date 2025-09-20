package update

import (
	"encoding/json"
	"fmt"
	m "ports/db"
	"time"
)

func (u *Update)UpHistoryChat(chatID int64, date time.Time, chat m.Chat) error {
	jsonBytes, _ := json.Marshal(chat)
	wrapped := "[" + string(jsonBytes) + "]"
	query := fmt.Sprintf(`
		UPDATE %s
		SET chat = jsonb_set(
			chat,
			'{history}',
			(chat->'history') || ($3::jsonb)
		)
		WHERE chat_id = $1
		AND date = $2;
	`, tabelChats)
	result, err := u.DB.Exec(query, chatID, date, wrapped)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("Error not row affected")
	}
	return nil
}