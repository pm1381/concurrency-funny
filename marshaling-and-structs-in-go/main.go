package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

type KeyboardButton struct {
	Text string `json:"text"`
	RequestContact bool `json:"request_contact"`
	RequestLocation bool `json:"request_location"`
}

type InlineKeyboardButton struct {
	Text string `json:"text"`
	CallbackData string `json:"callback_data"`
	Url string `json:"url"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	Keyboard [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard bool `json:"resize_keyboard"`
	OnTimeKeyboard bool `json:"one_time_keyboard"`
	Selective bool `json:"selective"`
}


type SendMessage struct {
	ChatID           interface{} `json:"chat_id"`
	Text             string `json:"text"`
	ParseMode        string `json:"parse_mode"`
	ReplyMarkup      interface{} `json:"reply_markup"`
}


func main()  {
	_, err := ReadSendMessageRequest("input_sample2.json")
	if err != nil {
		fmt.Println(err)
	}
}


func ReadSendMessageRequest(fileName string) (*SendMessage, error) {
	var message SendMessage
	var replyMarkup ReplyMarkup
	message.ReplyMarkup = replyMarkup
	readFromFile(fileName, &message)
	if message.ChatID == nil {return nil, errors.New("chat_id is empty")}
	if message.Text == "" {return nil, errors.New("text is empty")}
	if message.ReplyMarkup != nil {
		repMarkup := message.ReplyMarkup.(map[string]interface{})
		manageMarkup(repMarkup, &replyMarkup)
		message.ReplyMarkup = replyMarkup
	}
	fmt.Println(reflect.TypeOf(message.ReplyMarkup))
	fmt.Println(reflect.TypeOf(message.ReplyMarkup.(ReplyMarkup).InlineKeyboard))
	fmt.Println(len(message.ReplyMarkup.(ReplyMarkup).InlineKeyboard))
	fmt.Println(len(message.ReplyMarkup.(ReplyMarkup).InlineKeyboard[0]))
	fmt.Println((message.ReplyMarkup.(ReplyMarkup).InlineKeyboard[0][1].Text))
	fmt.Println((message.ParseMode))
	return &message, nil
}

func manageMarkup(replyM map[string]interface{}, replyMarkup *ReplyMarkup)  {
	_, ok := replyM["inline_keyboard"]
	var keyBoardData []interface{}
	if !ok {
		selective, ok := replyM["selective"]
		one_time_keyboard, ok2 := replyM["one_time_keyboard"]
		resize_keyboard, ok3 := replyM["resize_keyboard"]
		keyboard, ok4 := replyM["keyboard"]
		if (ok2 && ok && ok3 && ok4) {
			replyMarkup.Selective = selective.(bool)
			replyMarkup.OnTimeKeyboard = one_time_keyboard.(bool)
			replyMarkup.ResizeKeyboard = resize_keyboard.(bool)
			keyBoardData = (keyboard.([]interface{}))
		} else {
			log.Fatal("wrong input")
		}
	} else {
		keyBoardData = (replyM["inline_keyboard"].([]interface{}))
	}
	manageKeyboards(keyBoardData, replyMarkup)
}

func readFromFile(fileName string, messages *SendMessage) ()  {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(data), messages)
	if err != nil {
		log.Fatal(err)
	}
}

func manageKeyboards(iteration []interface{}, replyMarkup *ReplyMarkup) {
	var i, j int = 0, 0
	replyMarkup.InlineKeyboard = make([][]InlineKeyboardButton, len(iteration))
	replyMarkup.Keyboard = make([][]KeyboardButton, len(iteration))
	for _, val := range iteration {
		y := val.([]interface{})
		replyMarkup.Keyboard[i] = make([]KeyboardButton, len(y))
		replyMarkup.InlineKeyboard[i] = make([]InlineKeyboardButton, len(y))
		j = 0
		for _, v := range y {
			if reflect.TypeOf(v).Kind() == reflect.String {
				s := v
				if replyMarkup.InlineKeyboard == nil {
					replyMarkup.Keyboard[i][j] = KeyboardButton{
						Text: s.(string),
						RequestContact: s.(bool),
						RequestLocation: s.(bool),
					}
				} else {
					replyMarkup.InlineKeyboard[i][j] = InlineKeyboardButton{
						Text: s.(string),
						CallbackData: s.(string),
						Url: s.(string),
					}
				}
			} else {
				s := v.(map[string]interface{})
				if replyMarkup.InlineKeyboard == nil {
					var text, _ = s["text"].(string)
					var contact, _ = s["request_contact"].(bool)
					var location, _ = s["request_location"].(bool)
					replyMarkup.Keyboard[i][j] = KeyboardButton{
						Text: text,
						RequestContact: contact,
						RequestLocation: location,
					}
				} else {
					var text, _ = s["text"].(string)
					var callback, _ = s["callback_data"].(string)
					var url, _ = s["url"].(string)
					replyMarkup.InlineKeyboard[i][j] = InlineKeyboardButton{
						Text: text,
						CallbackData: callback,
						Url: url,
					}
				}
			}
			j++
		}
		i++
	}
}
