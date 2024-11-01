package main

import (
	"github.com/golang-concurrency-fun/internal/modules"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample1(t *testing.T) {
	b := modules.NewBaleImpl()
	_, err := b.AddUser("a", false)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid username", err.Error())
}

func TestSample2(t *testing.T) {
	b := modules.NewBaleImpl()
	id, err := b.AddUser("ali2000", false)
	assert.Nil(t, err)
	chatId, err := b.AddChat("quera", false, id, []int{id})
	assert.Nil(t, err)
	_, err = b.SendMessage(id, chatId, "salam")
	assert.Nil(t, err)
}
