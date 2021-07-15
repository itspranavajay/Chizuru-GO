package database

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"strings"
)

type User struct {
	UserId   int64    `gorm:"primary_key" json:"user_id"`
	UserName string `json:"user_name"`
}

type Chat struct {
	ChatId   string `gorm:"primary_key" json:"chat_id"`
	ChatName string `json:"chat_name"`
}

func EnsureBotInDb(b *gotgbot.Bot) {
	// Insert bot user only if it doesn't exist already
	botUser := &User{UserId: b.User.Id, UserName: b.User.Username}
	SESSION.Save(botUser)
}

func UpdateUser(userId int64, username string, chatId string, chatName string) {
	username = strings.ToLower(username)
	tx := SESSION.Begin()

	// upsert user
	user := &User{UserId: userId, UserName: username}
	tx.Save(user)

	if chatId == "nil" || chatName == "nil" {
		tx.Commit()
		return
	}

	// upsert chat
	chat := &Chat{ChatId: chatId, ChatName: chatName}
	tx.Save(chat)
	tx.Commit()
}

func GetStats()(users int64, chats int64){
	var UserCount int64
	SESSION.Model(&User{}).Count(&UserCount)

	var ChatCount int64
	SESSION.Model(&Chat{}).Count(&ChatCount)

	return UserCount, ChatCount

}