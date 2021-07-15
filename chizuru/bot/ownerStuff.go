package bot

import (
	"Chizuru-GO/chizuru/config"
	"Chizuru-GO/chizuru/database"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"runtime"
)

func getStats(b *gotgbot.Bot, ctx *ext.Context)(err error){
	if ctx.EffectiveMessage.From.Id != config.OwnerID {
		_, err = ctx.EffectiveMessage.Reply(b, "You're not allowed to use this command.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		return err
	}
	users, chats := database.GetStats()
	msg := fmt.Sprintf(`<b>%s, the lewd one near you.</b>
A project by <a href="https://github.com/Dank-del">Dank-del</a>
<b>Total users registered</b>: <code>%d</code>
<b>Total chats registered</b>: <code>%d</code>

<b>Go version</b>: <code>%s</code>`, b.User.FirstName, users, chats, runtime.Version())
	_, err = ctx.EffectiveMessage.Reply(b, msg, &gotgbot.SendMessageOpts{ParseMode: "html", DisableWebPagePreview: true})
    return err
}
