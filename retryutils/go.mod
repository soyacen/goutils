module github.com/soyacen/goutils/retryutils

go 1.15

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/soyacen/goutils/backoffutils v0.0.0-20211110092012-2d5c10c9a5dd
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/otel v1.3.0 // indirect
)

replace github.com/soyacen/goutils/backoffutils => ../backoffutils
