package bot

import (
	"Chizuru-GO/chizuru/nekosAPI"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var NekosHelpStr = fmt.Sprintf(`
<b>This is the help for the nekos.life package</b>

This bot supports a variety of endpoints, use them with the <code>/nekos</code> cmd.

Example: <code>/nekos neko</code>.

Supported endpoints:
<code>%s</code>
`, nekosAPI.EndPoints)

func getHelp(b *gotgbot.Bot, ctx *ext.Context) (err error) {
	if ctx.EffectiveChat.Type == "private" {
		_, err = ctx.EffectiveMessage.Reply(b, NekosHelpStr, &gotgbot.SendMessageOpts{ParseMode: "html"})
		return err
	} else {
		_, err = ctx.EffectiveMessage.Reply(b, "Command restricted to PM.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		return err
	}
}
