package ioutils

import "io"

// CloseThrowError close the closer and throw the close's error
func CloseThrowError(closer io.Closer, err *error) {
	if closer == nil {
		return
	}
	e := closer.Close()
	if e != nil && err != nil {
		*err = e
	}
}

// CloseQuietly close the closer and ignore the close's error
func CloseQuietly(closer io.Closer) {
	if closer == nil {
		return
	}
	closer.Close()
}
