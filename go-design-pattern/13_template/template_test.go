package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sms_Send(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", 1239999)
	assert.NoError(t, err)
}
