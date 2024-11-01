package modules

import (
	"errors"
	"sort"
)

type Chat struct {
	id        int
	name      string
	isGroup   bool
	isChannel bool
	creator   int
	admins    []int // based on MVC structure we can have []User also
}

func (b *messengerImpl) chatAuthorization(chat *Chat, user *User) bool {
	if chat.isChannel {
		var isAdmin bool = false
		for _, eachAdmin := range chat.admins {
			if eachAdmin == user.id {
				isAdmin = true
				break
			}
		}
		return isAdmin
	}
	return true
}

func (b *messengerImpl) findChat(chatId int) *Chat {
	for _, v := range b.chats {
		if v.id == chatId {
			return &v
		}
	}
	return nil
}

func (b *messengerImpl) BaleChatValidation(creator int) (string, bool) {
	for _, v := range b.users {
		if v.id == creator && v.isBot {
			return "could not create chat", true
		}
	}
	return "no error", false
}

func (b *messengerImpl) AddChat(chatname string, isGroup bool, creator int, admins []int) (int, error) {
	message, validation := b.BaleChatValidation(creator)
	if validation {
		return 0, errors.New(message)
	}
	newId := len(b.chats) + 1
	chat := Chat{
		id:        newId,
		name:      chatname,
		isGroup:   isGroup,
		isChannel: !isGroup,
		creator:   creator,
		admins:    admins,
	}
	b.chats = append(b.chats, chat)
	return newId, nil
}

func (b *messengerImpl) SetChatAdmin(chatId int, userId int) error {
	chat := b.findChat(chatId)
	isUserAdmin := sort.SearchInts(chat.admins, userId)
	if isUserAdmin >= 0 {
		return errors.New("user is already admin")
	}
	chat.admins = append(chat.admins, userId)
	return nil
}
