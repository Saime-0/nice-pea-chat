package service

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/saime-0/nice-pea-chat/internal/domain"
	"github.com/saime-0/nice-pea-chat/internal/repository/sqlite/memory"
	"github.com/stretchr/testify/assert"
)

func TestInvations_ChatInvitations(t *testing.T) {
	newInvitationsService := func() *Invitations {
		sqLiteInMemory, err := memory.Init(memory.Config{MigrationsDir: "../../migrations/repository/sqlite/memory"})
		assert.NoError(t, err)
		chatsRepository, err := sqLiteInMemory.NewChatsRepository()
		assert.NoError(t, err)
		membersRepository, err := sqLiteInMemory.NewMembersRepository()
		assert.NoError(t, err)
		invitationsRepository, err := sqLiteInMemory.NewInvitationsRepository()
		assert.NoError(t, err)

		return &Invitations{
			ChatsRepo:       chatsRepository,
			MembersRepo:     membersRepository,
			InvitationsRepo: invitationsRepository,
			History:         HistoryDummy{},
		}
	}
	t.Run("UserID обязательное поле", func(t *testing.T) {
		input := ChatInvitationsInput{
			UserID: "",
			ChatID: uuid.NewString(),
		}
		invsChat, err := newInvitationsService().ChatInvitations(input)
		assert.Error(t, err)
		assert.Len(t, invsChat, 0)
	})
	t.Run("ChatID обязательное поле", func(t *testing.T) {
		input := ChatInvitationsInput{
			ChatID: "",
			UserID: uuid.NewString(),
		}
		invsChat, err := newInvitationsService().ChatInvitations(input)
		assert.Error(t, err)
		assert.Len(t, invsChat, 0)
	})

	t.Run("UserID должен быть `chief` в чате с ChatID", func(t *testing.T) {
		newService := newInvitationsService()
		chatID := uuid.NewString()

		err := newService.ChatsRepo.Save(domain.Chat{
			ID:          chatID,
			Name:        "Name1",
			ChiefUserID: uuid.NewString(),
		})
		assert.NoError(t, err)
		input := ChatInvitationsInput{
			ChatID: chatID,
			UserID: uuid.NewString(),
		}
		invsChat, err := newService.ChatInvitations(input)
		assert.Error(t, err)
		assert.Len(t, invsChat, 0)
	})

	t.Run("пустой список из чата без приглошений", func(t *testing.T) {
		newService := newInvitationsService()
		chatID := uuid.NewString()
		userID := uuid.NewString()
		err := newService.ChatsRepo.Save(domain.Chat{
			ID:          chatID,
			Name:        "Name1",
			ChiefUserID: userID,
		})
		assert.NoError(t, err)
		input := ChatInvitationsInput{
			UserID: userID,
			ChatID: chatID,
		}
		invsChat, err := newService.ChatInvitations(input)
		assert.NoError(t, err)
		assert.Len(t, invsChat, 0)
	})

	t.Run("список из 4 приглашений из заполненного репозитория", func(t *testing.T) {
		newService := newInvitationsService()
		chatID := uuid.NewString()
		userID := uuid.NewString()
		err := newService.ChatsRepo.Save(domain.Chat{
			ID:          chatID,
			Name:        "Name1",
			ChiefUserID: userID,
		})
		assert.NoError(t, err)
		input := ChatInvitationsInput{
			UserID: userID,
			ChatID: chatID,
		}

		err = errors.Join(newService.InvitationsRepo.Save(domain.Invitation{
			ID:     uuid.NewString(),
			ChatID: chatID,
		}), newService.InvitationsRepo.Save(domain.Invitation{
			ID:     uuid.NewString(),
			ChatID: chatID,
		}), newService.InvitationsRepo.Save(domain.Invitation{
			ID:     uuid.NewString(),
			ChatID: chatID,
		}), newService.InvitationsRepo.Save(domain.Invitation{
			ID:     uuid.NewString(),
			ChatID: chatID,
		}))
		assert.NoError(t, err)
		invsChat, err := newService.ChatInvitations(input)
		assert.NoError(t, err)
		assert.Len(t, invsChat, 4)
	})
}
