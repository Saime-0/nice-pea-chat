package model

const (
	ControlConversation = 1
	ControlRoles        = 2
	Invite              = 3
	Kick                = 4
	SendMessages        = 5
	ControlMessages     = 6
	ReadHistory         = 7
)

var Permissions = []uint8{
	ControlConversation,
	ControlRoles,
	Invite,
	Kick,
	SendMessages,
	ControlMessages,
	ReadHistory,
}

type Permission struct {
	ID   uint8  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
