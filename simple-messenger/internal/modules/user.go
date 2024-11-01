package modules

import (
	"errors"
	"unicode"
)

type User struct {
	id       int // but we can use bigger. or also unit is correct
	isBot    bool
	username string
}

func (b *messengerImpl) findUser(userId int) (user *User) {
	for _, v := range b.users {
		if v.id == userId {
			return &v
		}
	}
	return nil
}

func (b *messengerImpl) BaleUserValidation(username string) (string, bool) {
	if len(username) < 3 {
		return "invalid username", true
	}
	hasLetter := false
	hasDigit := false
	for _, char := range username {
		if unicode.IsLetter(char) {
			hasLetter = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}
	}
	if !hasLetter || !hasDigit {
		return "invalid username", true
	}
	for _, v := range b.users {
		if v.username == username {
			return "invalid username", true
		}
	}
	return "no error", false
}

func (b *messengerImpl) AddUser(username string, isBot bool) (int, error) {
	message, validation := b.BaleUserValidation(username)
	if validation {
		return 0, errors.New(message)
	}
	newId := len(b.users) + 1
	user := User{
		username: username,
		isBot:    isBot,
		id:       newId,
	}
	b.users = append(b.users, user)
	return newId, nil
}

func (b *messengerImpl) GetLastUserMessage(userId int) (string, int, error) {
	var lastMessageID int
	var lastMessageText string
	for i := len(b.messages) - 1; i >= 0; i-- {
		message := b.messages[i]
		if message.senderId == userId {
			lastMessageID = message.id
			lastMessageText = message.text
			break
		}
	}
	return lastMessageText, lastMessageID, nil
}
