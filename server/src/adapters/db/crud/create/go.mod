module crud/create

go 1.22.1

replace ports/db => ../../../../ports/db

require (
	github.com/lib/pq v1.10.9
	ports/db v0.0.0-00010101000000-000000000000
)
