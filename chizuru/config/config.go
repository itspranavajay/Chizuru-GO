package config

import (
	"github.com/bigkevmcd/go-configparser"
	"log"
)


var ChatLINK string
var ChannelLINK string


func SetVars(p *configparser.ConfigParser){
	var err error
	ChatLINK, err = p.Get("telegram_links", "chat")
	if err != nil {
		log.Fatal("an error in getting chat link: ", err)
	}
	ChannelLINK, err = p.Get("telegram_links", "channel")
	if err != nil {
		log.Fatal("an error in getting channel link: ", err)
	}
}
