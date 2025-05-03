module github.com/glamostoffer/arete/learning

go 1.24.0

replace github.com/glamostoffer/arete/learning/pkg => ./pkg

require (
	github.com/glamostoffer/arete/learning/pkg v0.0.0-00010101000000-000000000000
	github.com/glamostoffer/arete/pkg v0.0.0-20250503135817-74ab4daf5b89
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.72.0
)

require (
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
