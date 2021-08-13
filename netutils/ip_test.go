package netutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicIP(t *testing.T) {
	ip, err := PublicIP()
	assert.Nil(t, err)
	assert.Equal(t, "172.16.40.45", ip)
}
