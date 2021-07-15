package config

import (
	"github.com/bigkevmcd/go-configparser"
	"log"
)

var ChatLINK string
var ChannelLINK string
var DBUrl string
var OwnerID int64

func SetVars(p *configparser.ConfigParser) {
	var err error
	DBUrl, err = p.Get("postgres", "url")
	if err != nil {
		log.Fatal("an error in getting postgresql database connection url: ", err)
	}

	OwnerID, err = p.GetInt64("bot_admins", "OwnerID")
	if err != nil {
		log.Fatal("an error in getting owner id: ", err)
	}

	ChatLINK, err = p.Get("telegram_links", "chat")
	if err != nil {
		log.Fatal("an error in getting chat link: ", err)
	}
	ChannelLINK, err = p.Get("telegram_links", "channel")
	if err != nil {
		log.Fatal("an error in getting channel link: ", err)
	}

}
