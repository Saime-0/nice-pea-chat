package app

import "github.com/saime-0/nice-pea-chat/internal/service"

type services struct {
	chats            *service.Chats
	invitations      *service.Invitations
	members          *service.Members
	sessions         *service.Sessions
	loginCredentials *service.LoginCredentials
}

func initServices(repos *repositories) *services {
	return &services{
		chats: &service.Chats{
			ChatsRepo:   repos.chats,
			MembersRepo: repos.members,
		},
		invitations: &service.Invitations{
			ChatsRepo:       repos.chats,
			MembersRepo:     repos.members,
			InvitationsRepo: repos.invitations,
			UsersRepo:       repos.users,
		},
		members: &service.Members{
			MembersRepo: repos.members,
			ChatsRepo:   repos.chats,
		},
		sessions: &service.Sessions{
			SessionsRepo: repos.sessions,
		},
		loginCredentials: &service.LoginCredentials{
			LoginCredentialsRepo: repos.loginCreds,
			SessionsRepo:         repos.sessions,
		},
	}
}
