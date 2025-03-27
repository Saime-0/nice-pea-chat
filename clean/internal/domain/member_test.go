package domain

import (
	"testing"
)

func TestMember_ValidateID(t *testing.T) {
	RunValidateIDTest(t, func(ID string) error {
		m := Member{ID: ID}
		return m.ValidateID()
	})
}

func TestMember_ValidateChatID(t *testing.T) {
	RunValidateChatIDTest(t, func(ChatID string) error {
		m := Member{ChatID: ChatID}
		return m.ValidateChatID()
	})
}
