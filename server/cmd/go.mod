module main

go 1.22.1

replace utils => ../utils

replace ports/db => ../src/ports/db

replace services/jwt => ../src/core/services/jwt

// web socket
replace ws/connection => ../ws/connection

replace ws/models => ../ws/models

replace ws/master => ../ws/master

//

// api
//     - routes user
replace routes/user => ../api/routes/user

replace user/login => ../api/routes/user/login

//     - routes server

//

// adapter-db
replace adapters/db => ../src/adapters/db

replace db/connection => ../src/adapters/db/connection

// replace db/models => ../src/adapters/db/models

replace crud/create => ../src/adapters/db/crud/create

replace crud/delete => ../src/adapters/db/crud/delete

replace crud/update => ../src/adapters/db/crud/update

replace crud/read => ../src/adapters/db/crud/read

//

require (
	adapters/db v0.0.0-00010101000000-000000000000
	db/connection v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.10.1
	ports/db v0.0.0-00010101000000-000000000000
	routes/user v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

require (
	crud/create v0.0.0-00010101000000-000000000000 // indirect
	crud/delete v0.0.0-00010101000000-000000000000 // indirect
	crud/read v0.0.0-00010101000000-000000000000 // indirect
	crud/update v0.0.0-00010101000000-000000000000 // indirect
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.3 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	services/jwt v0.0.0-00010101000000-000000000000 // indirect
	user/login v0.0.0-00010101000000-000000000000 // indirect
)
