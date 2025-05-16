module github.com/glamostoffer/arete/practice

go 1.24.0

replace github.com/glamostoffer/arete/practice/pkg => ./pkg

require (
	github.com/glamostoffer/arete/pkg v0.0.0-20250512193201-608b70ef9fde
	github.com/glamostoffer/arete/practice/pkg v0.0.0-00010101000000-000000000000
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	github.com/redis/go-redis/v9 v9.8.0
	github.com/tidwall/gjson v1.18.0
	github.com/tidwall/sjson v1.2.5
	google.golang.org/grpc v1.72.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/segmentio/kafka-go v0.4.47 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
