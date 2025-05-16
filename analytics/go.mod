module github.com/glamostoffer/arete/analytics

go 1.24.0

replace github.com/glamostoffer/arete/analytics/pkg => ./pkg

require (
	github.com/glamostoffer/arete/analytics/pkg v0.0.0-00010101000000-000000000000
	github.com/glamostoffer/arete/pkg v0.0.0-20250516162325-b094c2198cb7
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gofiber/fiber/v2 v2.52.6
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/jmoiron/sqlx v1.4.0
	github.com/segmentio/kafka-go v0.4.48
)

require (
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/grpc v1.71.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
