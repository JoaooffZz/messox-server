package db

import (
	m "db/models"
	"time"
)

type MockDB struct{}

func (mock *MockDB)GetNewUsers() []m.User {
	users := []m.User{
		{ID: 1, Name: "Alice", Password: "123456", Profile: "Developer", Bio: "Gopher apaixonada"},
		{ID: 2, Name: "Bob", Password: "abcdef", Profile: "Designer", Bio: "Ama UX/UI"},
		{ID: 3, Name: "Carol", Password: "qwerty", Profile: "Manager", Bio: "Organizada e focada"},
		{ID: 4, Name: "Dave", Password: "zxcvbn", Profile: "Tester", Bio: "Curioso e detalhista"},
		{ID: 5, Name: "Anna", Password: "pass123", Profile: "Analyst", Bio: "Boa em insights"},
		{ID: 6, Name: "Anselmo", Password: "pass234", Profile: "Engineer", Bio: "Adora resolver problemas"},
		{ID: 7, Name: "Chris", Password: "pass345", Profile: "Support", Bio: "Atencioso e prestativo"},
		{ID: 8, Name: "Christian", Password: "pass456", Profile: "DevOps", Bio: "Focado em automação"},
	}

	return users
}

func (mock *MockDB)GetNewChatData(chatID int64) m.ChatData {
	d, _ := time.Parse("2006-01-02", "2024-10-01")
	t, _ := time.Parse("15:04", "14:30")
	return m.ChatData{
		ChatID: chatID,
		Date: d,
		Chat: m.ChatJsonB{
			History: []m.Chat{
				{ID:0, Message: "Test Hello!", Time: t.Format("15:04")},
			},
		},
	}
}

func (mock *MockDB)GetNewInboxMessage(senderID int, addresseeID int) m.InboxMessage {
	t, _ := time.Parse("15:04", "15:30")
	return m.InboxMessage{
		SenderID: senderID,
		AddresseeID: addresseeID,
		Message: "Test inbox message",
		CreatedAt: t,
	}
}

func (mock *MockDB)GetNewChat(id int) (m.Chat, time.Time) {
	d, _ := time.Parse("2006-01-02", "2024-10-01")
	t, _ := time.Parse("15:04", "14:30")
	return m.Chat{
		ID: id,
		Message: "Teste update",
		Time: t.Format("15:04"),
	}, d
}