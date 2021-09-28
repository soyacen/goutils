module github.com/soyacen/goutils/retryutils

go 1.15

require (
	github.com/soyacen/goutils/backoffutils v0.0.0-20210915082908-e5fc170a08b3
	github.com/stretchr/testify v1.7.0
)

replace github.com/soyacen/goutils/backoffutils v0.0.0-20210915082908-e5fc170a08b3 => ../backoffutils
