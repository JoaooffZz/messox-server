package db

import (
	crudUp "crud/update"
	"database/sql"
	modelsDB "db/models"
	"fmt"
	chatID "services/chat_id"
)

type TestUpdateDB struct {
	DB    *sql.DB
	Users []modelsDB.User
}

func (t *TestUpdateDB) FlowRun() error {
	mock := MockDB{}
	update := crudUp.Update{DB: t.DB}

	// update profile
	if err := t.UnitUpUserProfile(update); err != nil {
		return fmt.Errorf("error in UnitUpUserProfile: %w", err)
	}

	// update bio
	if err := t.UnitUpUserBio(update); err != nil {
		return fmt.Errorf("error in UnitUpUserBio: %w", err)
	}

	// update history chat
	if err := t.UnitUpHistoryChat(mock, update); err != nil {
		return fmt.Errorf("error in UnitUpHistoryChat: %w", err)
	}

	return nil
}

func (t *TestUpdateDB) UnitUpUserProfile(update crudUp.Update) error {
	fmt.Printf("🔷 init UpUserProfile\n")
	for _, user := range t.Users {
		if err := update.UpUserProfile(user.ID, "test-update-profile"); err != nil {
			return fmt.Errorf("updating user profile (id %d): %w", user.ID, err)
		}
	}
	return nil
}

func (t *TestUpdateDB) UnitUpUserBio(update crudUp.Update) error {
	fmt.Printf("🔷 init UpUserBio\n")
	for _, user := range t.Users {
		if err := update.UpUserBio(user.ID, "test-update-bio"); err != nil {
			return fmt.Errorf("updating user bio (id %d): %w", user.ID, err)
		}
	}
	return nil
}

func (t *TestUpdateDB) UnitUpHistoryChat(mock MockDB, update crudUp.Update) error {
	fmt.Printf("🔷 init UpHistoryChat\n")
	for i := 0; i < len(t.Users); i++ {
		j := i + 1

		cID := chatID.BuildChatID(t.Users[i].ID, t.Users[j].ID)
		chat, d := mock.GetNewChat(t.Users[i].ID)

		if err := update.UpHistoryChat(cID, d, chat); err != nil {
			return fmt.Errorf("updating history chat (cid %s, user %d): %w", cID, t.Users[i].ID, err)
		}

		if j == len(t.Users)-1 {
			break
		}
		i = j
	}
	return nil
}
