package base

import (
	"testing"

	"github.com/saime-0/nice-pea-chat/internal/domain"
	"github.com/saime-0/nice-pea-chat/internal/domain/repository_tests"
)

func TestChatsRepository(t *testing.T) {
	repository_tests.ChatsRepositoryTests(t, func() domain.ChatsRepository {
		// todo: написать инициализацию базового репозитория, но в тестовом окружении
		return &ChatsRepository{}
	})
}
