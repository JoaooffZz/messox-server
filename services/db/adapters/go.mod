module adapters/db

go 1.22.1

replace ports/db => ../../ports/db

replace crud/create => ./crud/create

replace crud/delete => ./crud/delete

replace crud/update => ./crud/update

replace crud/read => ./crud/read

require (
	crud/create v0.0.0-00010101000000-000000000000
	crud/delete v0.0.0-00010101000000-000000000000
	crud/read v0.0.0-00010101000000-000000000000
	crud/update v0.0.0-00010101000000-000000000000
)

require ports/db v0.0.0-00010101000000-000000000000 // indirect
