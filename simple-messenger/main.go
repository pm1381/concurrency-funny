package main

import (
	"fmt"
	"github.com/golang-concurrency-fun/internal/modules"
)

func main() {
	b := modules.NewBaleImpl()
	id, err := b.AddUser("ali2000", false)
	fmt.Println(id)
	fmt.Println(err)
	chatId, err2 := b.AddChat("quera", false, id, []int{id})
	fmt.Println(chatId)
	fmt.Println(err2)
	sendMessageId, err3 := b.SendMessage(id, chatId, "salam")
	fmt.Println(sendMessageId)
	fmt.Println(err3)
}
