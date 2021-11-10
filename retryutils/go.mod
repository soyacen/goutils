module github.com/soyacen/goutils/retryutils

go 1.15

require (
	github.com/soyacen/goutils/backoffutils v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
)

replace github.com/soyacen/goutils/backoffutils => ../backoffutils
