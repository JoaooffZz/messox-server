package db

import (
	c "crud/create"
	d "crud/delete"
	r "crud/read"
	u "crud/update"
	"database/sql"
)

type AdapterDB struct {
	c.Create
	r.Read
	u.Update
	d.Delete
}

func New(db *sql.DB) *AdapterDB {
	return &AdapterDB{
		c.Create{DB: db},
		r.Read{DB: db},
		u.Update{DB: db},
		d.Delete{DB: db},
	}
}