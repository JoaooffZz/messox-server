package db

import (
	crudRead "crud/read"
	"database/sql"
	modelsDB "db/models"
	"encoding/json"
	"fmt"
)

type TestReadDB struct {
	DB    *sql.DB
	Users []modelsDB.User
}

func (t *TestReadDB) FlowRun() error {
	read := crudRead.Read{DB: t.DB}

	// total users
	if err := t.UnitGetTotalUsers(read); err != nil {
		return fmt.Errorf("error in UnitGetTotalUsers: %w", err)
	}

	// users by char
	if err := t.UnitGetUsers(read); err != nil {
		return fmt.Errorf("error in UnitGetUsers: %w", err)
	}

	// inbox messages
	if err := t.UnitGetInboxMessages(read); err != nil {
		return fmt.Errorf("error in UnitGetInboxMessages: %w", err)
	}

	// sent requests
	if err := t.UnitGetSentRequests(read); err != nil {
		return fmt.Errorf("error in UnitGetSentRequests: %w", err)
	}

	// received requests
	if err := t.UnitGetReceivedRequests(read); err != nil {
		return fmt.Errorf("error in UnitGetReceivedRequests: %w", err)
	}

	return nil
}

func (t *TestReadDB) UnitGetTotalUsers(read crudRead.Read) error {
	fmt.Printf("🔷 init GetTotalUsers\n")
	total, err := read.GetTotalUsers()
	if err != nil {
		return fmt.Errorf("get total users: %w", err)
	}
	if total != len(t.Users) {
		return fmt.Errorf("total users inconsistency [total: %v][users-list: %v]", total, len(t.Users))
	}
	return nil
}

func (t *TestReadDB) UnitGetUsers(read crudRead.Read) error {
	fmt.Printf("🔷 init GetUsers\n")
	for _, user := range t.Users {
		char := user.Name[:1]
		usersView, err := read.GetUsers(char)
		if err != nil {
			return fmt.Errorf("get users by char %q: %w", char, err)
		}

		fmt.Printf("\nchar: %v", char)
		for _, view := range usersView {
			jsonData, err := json.MarshalIndent(view, "", "  ")
			if err != nil {
				return fmt.Errorf("marshal user view: %w", err)
			}
			fmt.Println(string(jsonData))
		}
	}
	return nil
}

func (t *TestReadDB) UnitGetInboxMessages(read crudRead.Read) error {
	fmt.Printf("🔷 init GetInboxMessages\n")
	for _, user := range t.Users {
		fmt.Printf("\nuser: %v", user)
		inbox, err := read.GetInboxMessages(user.ID)
		if err != nil {
			return fmt.Errorf("get inbox messages for user %d: %w", user.ID, err)
		}
		jsonInbox, err := json.Marshal(inbox)
		if err != nil {
			return fmt.Errorf("marshal inbox messages: %w", err)
		}
		fmt.Println(string(jsonInbox))
	}
	return nil
}

func (t *TestReadDB) UnitGetSentRequests(read crudRead.Read) error {
	fmt.Printf("🔷 init GetSentRequests\n")
	for _, user := range t.Users {
		fmt.Printf("\nuser: %v", user.Name)
		usersView, err := read.GetSentRequests(user.ID)
		if err != nil {
			return fmt.Errorf("get sent requests for user %d: %w", user.ID, err)
		}
		fmt.Printf("\nsents total: %v", len(usersView))
		for _, sent := range usersView {
			jsonData, err := json.MarshalIndent(sent, "", "  ")
			if err != nil {
				return fmt.Errorf("marshal sent request: %w", err)
			}
			fmt.Println(string(jsonData))
		}
	}
	return nil
}

func (t *TestReadDB) UnitGetReceivedRequests(read crudRead.Read) error {
	fmt.Printf("🔷 init GetReceivedRequests\n")
	for _, user := range t.Users {
		fmt.Printf("\nuser: %v", user.Name)
		usersView, err := read.GetReceivedRequests(user.ID)
		if err != nil {
			return fmt.Errorf("get received requests for user %d: %w", user.ID, err)
		}
		fmt.Printf("\nreceiveds total: %v", len(usersView))
		for _, received := range usersView {
			jsonData, err := json.MarshalIndent(received, "", "  ")
			if err != nil {
				return fmt.Errorf("marshal received request: %w", err)
			}
			fmt.Println(string(jsonData))
		}
	}
	return nil
}
