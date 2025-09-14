package update

import (
	m "db/models"
	"encoding/json"
	"fmt"
	"time"
)

func (u *Update)UpHistoryChat(chatID int64, date time.Time, chat m.Chat) error {
	// query := fmt.Sprintf(`
	// 	UPDATE %s
	// 	SET chat = jsonb_set(
	// 		chat,
	// 		'{history}',
	// 		COALESCE(chat->'history', '[]'::jsonb) || jsonb_build_array(
	// 			jsonb_build_object(
	// 				'id', $3::INT,
	// 				'message', $4::TEXT,
	// 				'time', $5::TIME
	// 			)
	// 		)
	// 	)
	// 	WHERE chat_id = $1
	// 	AND date = $2;
	// `, tabelChats)
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