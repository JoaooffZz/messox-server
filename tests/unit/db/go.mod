module test/db

go 1.22.1

// core services
replace services/chat_id => ../../server/src/core/services/chat_id

//

// adapters db
replace db/models => ../../server/src/adapters/db/models

replace crud/create => ../../server/src/adapters/db/crud/create

replace crud/read => ../../server/src/adapters/db/crud/read

replace crud/update => ../../server/src/adapters/db/crud/update

replace crud/delete => ../../server/src/adapters/db/crud/delete

//

require (
	crud/create v0.0.0-00010101000000-000000000000
	crud/delete v0.0.0-00010101000000-000000000000
	crud/read v0.0.0-00010101000000-000000000000
	crud/update v0.0.0-00010101000000-000000000000
	db/models v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v4 v4.18.3
	services/chat_id v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	golang.org/x/crypto v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
