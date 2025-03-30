package service

import (
	"errors"

	"github.com/google/uuid"

	"github.com/saime-0/nice-pea-chat/internal/domain"
)

type Members struct {
	MembersRepo domain.MembersRepository
	ChatsRepo   domain.ChatsRepository
}

type ChatMembersInput struct {
	SubjectUserID string
	ChatID        string
}

var (
	ErrChatMembersInputSubjectUserIDValidate = errors.New("некорректный SubjectUserID")
	ErrChatMembersInputChatIDValidate        = errors.New("некорректный ChatID")
)

// Validate валидирует значение отдельно каждого параметры
func (in ChatMembersInput) Validate() error {
	if err := uuid.Validate(in.SubjectUserID); err != nil {
		return errors.Join(err, ErrChatMembersInputSubjectUserIDValidate)
	}
	if err := uuid.Validate(in.ChatID); err != nil {
		return errors.Join(err, ErrChatMembersInputChatIDValidate)
	}

	return nil
}

var (
	ErrChatMembersChatNotExists   = errors.New("чата с таким ID не существует")
	ErrChatMembersUserIsNotMember = errors.New("пользователь не является участником чата")
)

// ChatMembers возвращает список участников чата
func (m *Members) ChatMembers(in ChatMembersInput) ([]domain.Member, error) {
	// Валидировать параметры
	var err error
	if err = in.Validate(); err != nil {
		return nil, err
	}

	// Проверить существование чата
	chatsFilter := domain.ChatsFilter{
		IDs: []string{in.ChatID},
	}
	chats, err := m.ChatsRepo.List(chatsFilter)
	if err != nil {
		return nil, err
	}
	if len(chats) != 1 {
		return nil, ErrChatMembersChatNotExists
	}

	// Получить список участников
	membersFilter := domain.MembersFilter{
		ChatID: in.ChatID,
	}
	members, err := m.MembersRepo.List(membersFilter)
	if err != nil {
		return nil, err
	}
	if len(members) == 0 {
		return nil, ErrChatMembersUserIsNotMember
	}

	// Проверить что пользователь является участником чата
	for i, member := range members {
		if member.UserID == in.SubjectUserID {
			break
		} else if i == len(members)-1 {
			return nil, ErrChatMembersUserIsNotMember
		}
	}

	return members, nil
}

type LeaveInput struct {
	SubjectUserID string
	ChatID        string
}

func (in LeaveInput) Validate() error {
	if err := uuid.Validate(in.SubjectUserID); err != nil {
		return errors.Join(err, ErrChatMembersInputSubjectUserIDValidate)
	}
	if err := uuid.Validate(in.ChatID); err != nil {
		return errors.Join(err, ErrChatMembersInputChatIDValidate)
	}

	return nil
}

var (
	ErrMembersLeaveChatNotExists        = errors.New("чата с таким ID не существует")
	ErrMembersLeaveUserIsNotMember      = errors.New("пользователь не является участником чата")
	ErrMembersLeaveUserShouldNotBeChief = errors.New("пользователь является главным администратором чата")
)

// Leave удаляет участника из чата
func (m *Members) Leave(in LeaveInput) error {
	// Валидировать параметры
	var err error
	if err = in.Validate(); err != nil {
		return err
	}

	// Проверить существование чата
	chatsFilter := domain.ChatsFilter{
		IDs: []string{in.ChatID},
	}
	chats, err := m.ChatsRepo.List(chatsFilter)
	if err != nil {
		return err
	}
	if len(chats) != 1 {
		return ErrMembersLeaveChatNotExists
	}

	// Пользователь должен быть участником чата
	membersFilter := domain.MembersFilter{
		UserID: in.SubjectUserID,
		ChatID: in.ChatID,
	}
	members, err := m.MembersRepo.List(membersFilter)
	if err != nil {
		return err
	}
	if len(members) != 1 {
		return ErrMembersLeaveUserIsNotMember
	}

	// Пользователь не должен быть главным администратором
	if members[0].UserID == chats[0].ChiefUserID {
		return ErrMembersLeaveUserShouldNotBeChief
	}

	// Удалить пользователя из чата
	if err = m.MembersRepo.Delete(members[0].ID); err != nil {
		return err
	}

	return nil
}
