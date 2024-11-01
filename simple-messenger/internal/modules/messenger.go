package modules

type Messenger interface {
	AddUser(username string, isBot bool) (int, error)
	AddChat(chatname string, isGroup bool, creator int, admins []int) (int, error)
	SendMessage(userId, chatId int, text string) (int, error)
	SendLike(userId, messageId int) error
	GetNumberOfLikes(messageId int) (int, error)
	SetChatAdmin(chatId, userId int) error
	GetLastMessage(chatId int) (string, int, error)
	GetLastUserMessage(userId int) (string, int, error)
}

type messengerImpl struct {
	users    []User
	chats    []Chat
	messages []Message
	likes    []Like
}

func NewBaleImpl() Messenger {
	return &messengerImpl{
		users:    make([]User, 0),
		chats:    make([]Chat, 0),
		messages: make([]Message, 0),
		likes:    make([]Like, 0),
	}
}
