package bot

import (
	"Chizuru-GO/chizuru/config"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func startCMD(b *gotgbot.Bot, ctx *ext.Context) (err error) {
	if ctx.Message.Chat.Type == "private" {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello, I'm %s, I fetch lewds just for you UwU\nTap on /help to know more!", b.User.FirstName), &gotgbot.SendMessageOpts{
			ParseMode: "html",
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Channel", Url: config.ChannelLINK},
					{Text: "Chat", Url: config.ChatLINK},
				}},
			},
		})
	} else {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hey there, *I'm %s*", b.User.FirstName), &gotgbot.SendMessageOpts{
			ParseMode: "markdownv2",
		})

	}
	if err != nil {
		fmt.Println("failed to send: " + err.Error())
	}
	return nil
}
