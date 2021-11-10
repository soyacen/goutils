module github.com/soyacen/goutils/retryutils

go 1.15

require (
	github.com/soyacen/goutils/backoffutils v0.0.0
	github.com/stretchr/testify v1.7.0
)

replace github.com/soyacen/goutils/backoffutils v0.0.0 => ../backoffutils
