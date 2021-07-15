package bot

import (
	"Chizuru-GO/chizuru/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func logStuff(b *gotgbot.Bot, ctx *ext.Context)(err error){
	u := ctx.EffectiveMessage.From
	c := ctx.EffectiveChat
	database.UpdateUser(u.Id, u.Username, string(c.Id), c.Title)
	return nil
}
