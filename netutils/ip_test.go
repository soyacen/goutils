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

func TestInterfaceIP(t *testing.T) {
	ip, err := InterfaceIP("en7")
	assert.NoError(t, err)
	assert.NotEmpty(t, ip)
	t.Log(ip)
}
