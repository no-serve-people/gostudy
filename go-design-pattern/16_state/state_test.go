package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMachine_GetStateName(t *testing.T) {
	m := &Machine{
		GetLeaderApproveState(),
	}
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Reject()
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Approval()
}
