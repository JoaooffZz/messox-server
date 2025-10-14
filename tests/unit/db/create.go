package db

import (
	crudCreate "crud/create"
	"database/sql"
	dbModels "db/models"
	"fmt"
	serviceChatID "services/chat_id"
)

type TestCreateDB struct {
	DB *sql.DB
	Users []dbModels.User
}

func (t *TestCreateDB) FlowCreate() error {
	mock := MockDB{}
	create := crudCreate.Create{DB: t.DB}

	// create users
	err := t.UnitCreateUser(create)
	if err != nil {
		return fmt.Errorf("error test in UnitCreateUser: %w", err)
	}

	// create contacts
	err = t.UnitCreateContacts(create)
	if err != nil {
		return fmt.Errorf("failed exec test in UnitCreateContacts: %w", err)
	}

	// create history chats
	err = t.UnitCreateHistoryChats(mock, create)
	if err != nil {
		return fmt.Errorf("failed exec test in UnitCreateHistoryChats: %w", err)
	}

	// create inbox messages
	err = t.UnitCreateInboxMessages(mock, create)
	if err != nil {
		return fmt.Errorf("failed exec test in UnitCreateInboxMessages: %w", err)
	}

	// create inbox requests in pairs
	err = t.UnitCreateInboxRequests(create)
	if err != nil {
		return fmt.Errorf("failed exec test in UnitCreateInboxRequests: %w", err)
	}

	return nil
}

func (t *TestCreateDB)authLenListUsers() error {
	if (len(t.Users) % 2) != 0  {
		return fmt.Errorf("for the tests to be executed it is necessary that the user list is in pairs")
	}
	return nil
}

func (t *TestCreateDB)UnitCreateUser(create crudCreate.Create) error {
	err := t.authLenListUsers()
	if err != nil {
		return err
	}
	for _, user := range t.Users {
		_, err := create.NewUser(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TestCreateDB)UnitCreateContacts(create crudCreate.Create) error {
	err := t.authLenListUsers()
	if err != nil {
		return err
	}
	for i := 0; i < len(t.Users); i++{
		j := i + 1
		err := create.NewContact(t.Users[i].ID, t.Users[j].ID)
		if err != nil {
			return err
		}
		if j == len(t.Users) - 1 {
			break
		}
		i = j
	}
	return nil
}

func (t *TestCreateDB)UnitCreateHistoryChats(mock MockDB, create crudCreate.Create) error {
	err := t.authLenListUsers()
	if err != nil {
		return err
	}
	for i := 0; i < len(t.Users); i++{
		j := i + 1

		chatID := serviceChatID.BuildChatID(t.Users[i].ID, t.Users[j].ID)
		chat := mock.GetNewChatData(chatID)
		err := create.NewHistoryChat(chat)
		if err != nil {
			return err
		}

		if j == len(t.Users) - 1 {
			break
		}
		i = j
	}
	return nil
}

func (t *TestCreateDB)UnitCreateInboxMessages(mock MockDB, create crudCreate.Create) error {
	err := t.authLenListUsers()
	if err != nil {
		return err
	}
	for i := 0; i < len(t.Users); i++{
		j := i + 1

		inbox := mock.GetNewInboxMessage(t.Users[i].ID, t.Users[j].ID)
		err := create.NewInboxMessage(inbox)
		if err != nil {
			return err
		}

		if j == len(t.Users) - 1 {
			break
		}
		i = j
	}
	return nil
}

func (t *TestCreateDB)UnitCreateInboxRequests(create crudCreate.Create) error {
	err := t.authLenListUsers()
	if err != nil {
		return err
	}
	for i := 0; i < len(t.Users); i++{
		j := i + 1
		err := create.NewInboxRequest(t.Users[i].ID, t.Users[j].ID)
				if err != nil {
			return err
		}

		if j == len(t.Users) - 1 {
			break
		}
		i = j
	}
	return nil
}