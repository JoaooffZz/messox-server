package db

import (
	crudDel "crud/delete"
	"database/sql"
	modelsDB "db/models"
	"fmt"
)

type TestDeleteDB struct {
	DB    *sql.DB
	Users []modelsDB.User
}

func (t *TestDeleteDB) FlowRun() error {
	del := crudDel.Delete{DB: t.DB}

	// Delete contacts
	if err := t.UnitDeleteContacts(del); err != nil {
		return fmt.Errorf("error in UnitDeleteContacts: %w", err)
	}

	// Delete inbox messages
	if err := t.UnitDeleteInboxMessages(del); err != nil {
		return fmt.Errorf("error in UnitDeleteInboxMessages: %w", err)
	}

	// Delete inbox requests
	if err := t.UnitDeleteInboxRequests(del); err != nil {
		return fmt.Errorf("error in UnitDeleteInboxRequests: %w", err)
	}

	return nil
}

func (t *TestDeleteDB) UnitDeleteContacts(del crudDel.Delete) error {
	fmt.Printf("🔷 init DeleteContacts\n")
	for i := 0; i < len(t.Users); i++ {
		j := i + 1

		if err := del.DeleteContact(t.Users[i].ID, t.Users[j].ID); err != nil {
			return fmt.Errorf("delete contact (%d, %d): %w", t.Users[i].ID, t.Users[j].ID, err)
		}

		if j == len(t.Users)-1 {
			break
		}
		i = j
	}
	return nil
}

func (t *TestDeleteDB) UnitDeleteInboxMessages(del crudDel.Delete) error {
	fmt.Printf("🔷 init DeleteInboxMessages\n")
	for i := 1; i < len(t.Users); i++ {
		if err := del.DeleteInboxMessages(t.Users[i].ID); err != nil {
			return fmt.Errorf("delete inbox messages (user %d): %w", t.Users[i].ID, err)
		}
		i++
	}
	return nil
}

func (t *TestDeleteDB) UnitDeleteInboxRequests(del crudDel.Delete) error {
	fmt.Printf("🔷 init DeleteInboxRequests\n")
	for i := 0; i < len(t.Users); i++ {
		j := i + 1

		if err := del.DeleteInboxRequest(t.Users[i].ID, t.Users[j].ID); err != nil {
			return fmt.Errorf("delete inbox request (%d, %d): %w", t.Users[i].ID, t.Users[j].ID, err)
		}

		if j == len(t.Users)-1 {
			break
		}
		i = j
	}
	return nil
}
