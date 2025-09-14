package models

import (
	"encoding/json"
	"time"
)

type User struct {
    ID       int    `db:"id" json:"id"`
    Name     string `db:"name" json:"name"`
    Password string `db:"password" json:"password"`
    Profile  string `db:"profile" json:"profile"`
    Bio      string `db:"bio" json:"bio"`
}

type ViewUser struct {
    Name    string `db:"name" json:"name"`
    Profile string `db:"profile" json:"profile"`
    Bio     string `db:"bio" json:"bio"`
}

type InboxMessage struct {
    SenderID    int       `db:"sender_id" json:"sender_id"`
    AddresseeID int       `db:"addressee_id" json:"addressee_id"`
    Message     string    `db:"message" json:"message"`
    CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
func (m InboxMessage) MarshalJSON() ([]byte, error) {
    type Alias InboxMessage // cria um tipo auxiliar para evitar recursão
    return json.Marshal(&struct {
        CreatedAt string `json:"created_at"` // campo customizado
        *Alias
    }{
        CreatedAt: m.CreatedAt.Format("15:04"), // formata só hora:minuto
        Alias:     (*Alias)(&m),                // copia todos os outros campos
    })
}

type ChatData struct {
    ChatID int64                  `db:"chat_id" json:"chat_id"`
    Date   time.Time              `db:"date" json:"date"`
    Chat   ChatJsonB `db:"chat" json:"chat"`
}
type ChatJsonB struct {
    History []Chat `db:"history" json:"history"`
}
type Chat struct {
    ID int `db:"id" json:"id"`
    Message string `db:"message" json:"message"`
    Time string `db:"time" json:"time"`
}