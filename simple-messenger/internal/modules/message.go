package modules

import (
	"errors"
)

type Message struct {
	id       int
	senderId int
	chatId   int
	text     string
}

func (b *messengerImpl) findMessage(messageId int) (message *Message) {
	for _, v := range b.messages {
		if v.id == messageId {
			return &v
		}
	}
	return nil
}

func (b *messengerImpl) BaleMessageValidation(chatId int, senderId int) (string, bool) {
	chat := findChat(chatId, b)
	user := b.findUser(senderId)
	if chat == nil {
		return "chat not found", true
	}
	if user == nil {
		return "user not found", true
	}
	if !chatAuthorization(chat, user) {
		return "user could not send message", true
	}
	return "no error", false
}

func (b *messengerImpl) SendMessage(senderId int, chatId int, text string) (int, error) {
	message, validation := b.BaleMessageValidation(chatId, senderId)
	if validation {
		return 0, errors.New(message)
	}
	newId := len(b.messages) + 1
	newMessage := Message{
		id:       newId,
		text:     text,
		senderId: senderId,
		chatId:   chatId,
	}
	b.messages = append(b.messages, newMessage)
	return newId, nil
}

func (b *messengerImpl) GetLastMessage(chatId int) (string, int, error) {
	var lastMessageID int
	var lastMessageText string
	for i := len(b.messages) - 1; i >= 0; i-- {
		message := b.messages[i]
		if message.chatId == chatId {
			lastMessageID = message.id
			lastMessageText = message.text
			break
		}
	}
	return lastMessageText, lastMessageID, nil
}
