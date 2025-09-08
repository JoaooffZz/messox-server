package db

type DB interface {
	CreateDB
	ReadDB
	UpdateDB
	DeleteDB
}