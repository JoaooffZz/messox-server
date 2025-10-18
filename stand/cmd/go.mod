module main

go 1.23.0

toolchain go1.24.7

replace utils => ../utils

replace ports/db => ../src/ports/db

//

// api
replace middleware/jwt => ../api/middleware/jwt

replace middleware/headers => ../api/middleware/headers

//     - routes user
replace routes/user => ../api/routes/user

replace user/login => ../api/routes/user/login

replace user/register => ../api/routes/user/register

//     - routes server
replace routes/server => ../api/routes/server

replace server/ping => ../api/routes/server/ping

replace server/ws => ../api/routes/server/web_socket

//

// adapter-db
replace adapters/db => ../src/adapters/db

replace db/connection => ../src/adapters/db/connection

replace crud/create => ../src/adapters/db/crud/create

replace crud/delete => ../src/adapters/db/crud/delete

replace crud/update => ../src/adapters/db/crud/update

replace crud/read => ../src/adapters/db/crud/read

//

// websocket-server
replace ws/connection => ../ws/connection

replace ws/models => ../ws/models

replace ws/master => ../ws/master

//

require (
	adapters/db v0.0.0-00010101000000-000000000000
	db/connection v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.11.0
	ports/db v0.0.0-00010101000000-000000000000
	routes/server v0.0.0-00010101000000-000000000000
	routes/user v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
	ws/connection v0.0.0-00010101000000-000000000000
)

require (
	crud/create v0.0.0-00010101000000-000000000000 // indirect
	crud/delete v0.0.0-00010101000000-000000000000 // indirect
	crud/read v0.0.0-00010101000000-000000000000 // indirect
	crud/update v0.0.0-00010101000000-000000000000 // indirect
	github.com/bytedance/sonic v1.14.0 // indirect
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.27.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/goccy/go-yaml v1.18.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/quic-go/quic-go v0.54.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.3.0 // indirect
	go.uber.org/mock v0.5.0 // indirect
	golang.org/x/arch v0.20.0 // indirect
	golang.org/x/crypto v0.40.0 // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	golang.org/x/tools v0.34.0 // indirect
	google.golang.org/protobuf v1.36.9 // indirect
	middleware/headers v0.0.0-00010101000000-000000000000 // indirect
	middleware/jwt v0.0.0-00010101000000-000000000000 // indirect
	server/ping v0.0.0-00010101000000-000000000000 // indirect
	server/ws v0.0.0-00010101000000-000000000000 // indirect
	user/login v0.0.0-00010101000000-000000000000 // indirect
	user/register v0.0.0-00010101000000-000000000000 // indirect
	ws/master v0.0.0-00010101000000-000000000000 // indirect
	ws/models v0.0.0-00010101000000-000000000000 // indirect
)
