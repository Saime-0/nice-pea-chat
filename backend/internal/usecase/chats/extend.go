package chats

import (
	"gorm.io/gorm"

	extend1 "github.com/saime-0/nice-pea-chat/internal/app/extend"
	"github.com/saime-0/nice-pea-chat/internal/model/rich"
)

type chatsExt struct {
	chatIDs []uint
	chats   map[uint]*rich.Chat
	db      *gorm.DB
}

func (e *chatsExt) lastMessages() (field extend1.Field) {
	field.Key = "last_message"
	field.Deps = nil
	field.Fn = func() error {
		lastMsgs := make([]rich.Message, 0, len(e.chats))
		if err := e.db.Raw(`
				SELECT DISTINCT ON (messages.chat_id) *
				FROM messages
					LEFT JOIN users AS author
						ON author.id = messages.author_id
					LEFT JOIN messages AS reply
						ON reply.id = messages.reply_to_id
					LEFT JOIN users AS reply_author
						ON reply_author.id = reply.author_id
				WHERE messages.chat_id IN (?) 
				ORDER BY messages.chat_id, messages.id DESC`,
			e.chatIDs,
		).Scan(&lastMsgs).Error; err != nil {
			return err
		}

		// Save into chatsExt
		for _, msg := range rich.MsgsMap(lastMsgs) {
			e.chats[msg.ChatID].LastMessage = msg
		}

		return nil
	}

	return field
}

func (e *chatsExt) unreadCounter(userID uint) (field extend1.Field) {
	field.Key = "unread_counter"
	field.Deps = nil
	field.Fn = func() error {
		//var unreadByChatID map[uint]int
		var unreads []struct {
			ChatID uint
			Count  int
		}
		if err := e.db.Raw(`
			SELECT DISTINCT ON (messages.chat_id) 
				messages.chat_id AS chat_id,
				count(messages.*)
			FROM messages
				INNER JOIN members mem
					ON mem.chat_id = messages.chat_id
			WHERE messages.id > coalesce(mem.last_read_msg_id, 0)
				AND messages.removed_at IS NULL
				AND mem.user_id = ?
			GROUP BY messages.chat_id`,
			userID,
		).Scan(&unreads).Error; err != nil {
			return err
		}

		// Save into chatsExt
		for _, unread := range unreads {
			e.chats[unread.ChatID].UnreadMessagesCount = unread.Count
		}

		return nil
	}

	return field
}

func extend(out *Out, p Params) error {
	ext := &chatsExt{
		db:      p.DB,
		chats:   make(map[uint]*rich.Chat, len(out.Chats)),
		chatIDs: make([]uint, len(out.Chats)),
	}

	// Fill required fields for extendParams
	for i, chat := range out.Chats {
		ext.chatIDs[i] = chat.ID
		ext.chats[chat.ID] = &chat
	}

	// Extend Params
	extendParams := (&extend1.Params{}).
		AddField(ext.lastMessages())

	// Add optional extenders
	if p.UnreadCounterForUser.IsSet {
		extendParams.AddField(ext.unreadCounter(p.UnreadCounterForUser.Val))
	}

	// Run extending
	if err := extendParams.Run(); err != nil {
		return err
	}

	out.Chats = make([]rich.Chat, 0, len(ext.chats))
	for _, chat := range ext.chats {
		out.Chats = append(out.Chats, *chat)
	}

	return nil
}
