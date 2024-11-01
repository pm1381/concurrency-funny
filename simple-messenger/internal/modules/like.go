package modules

import (
	"errors"
)

type Like struct {
	id      int
	message int
	user    int
	// also we could have a like property for each message but this is not
	// accurate when considering database because we will have multiple values
	// for a column
}

func (b *messengerImpl) GetNumberOfLikes(messageId int) (int, error) {
	count := 0
	for _, v := range b.likes {
		if v.message == messageId {
			count++
		}
	}
	return count, nil
}

func (b *messengerImpl) SendLike(userId int, messageId int) error {
	message, validation := b.BaleLikeValidation(userId, messageId)
	if validation {
		return errors.New(message)
	}
	newId := len(b.likes) + 1
	like := Like{
		id:      newId,
		message: messageId,
		user:    userId,
	}
	b.likes = append(b.likes, like)
	return nil
}

func (b *messengerImpl) BaleLikeValidation(userId int, messageId int) (string, bool) {
	message := b.findMessage(messageId)
	user := b.findUser(userId)
	if message == nil {
		return "message not found", true
	}
	if user == nil {
		return "user not found", true
	}
	for _, v := range b.likes {
		if v.user == userId && v.message == messageId {
			return "this user has liked this message before", true
		}
	}
	return "", false
}
