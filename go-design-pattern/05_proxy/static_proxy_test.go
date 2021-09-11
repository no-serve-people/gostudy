package proxy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserProxy_Login(t *testing.T) {
	proxy := NewUserProxy(&User{})
	err := proxy.Login("test", "password")

	require.Nil(t, err)
}
