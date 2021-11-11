module github.com/soyacen/goutils/retryutils

go 1.15

require (
	github.com/soyacen/goutils/backoffutils v0.0.0-20211110092012-2d5c10c9a5dd
	github.com/stretchr/testify v1.7.0
)

replace github.com/soyacen/goutils/backoffutils => ../backoffutils
