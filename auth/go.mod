module github.com/glamostoffer/arete/auth

go 1.24.0

replace github.com/glamostoffer/arete/auth/pkg => ./pkg

require (
	github.com/glamostoffer/arete/auth/pkg v0.0.0-00010101000000-000000000000
	github.com/glamostoffer/arete/pkg v0.0.0-20250408200434-b1a3f717d022
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/jmoiron/sqlx v1.4.0
	github.com/redis/go-redis/v9 v9.7.3
	golang.org/x/crypto v0.37.0
	google.golang.org/grpc v1.71.0
)

require github.com/lib/pq v1.10.9 // indirect

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
)
