package bot

import (
	"Chizuru-GO/chizuru/nekosAPI"
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func nekosLIFE(b *gotgbot.Bot, ctx *ext.Context) (err error) {
	args := ctx.Args()
	if len(args) == 1 {
		_, err = ctx.EffectiveMessage.Reply(b, "No argument provided.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		return err
	}

	sat := false
	for _, en := range nekosAPI.EndPoints {
		if en == args[1] {
			sat = true
			pic, err := nekosAPI.DoRequest(args[1])
			if err != nil {
				_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Error: <code>%s</code>", err.Error()), &gotgbot.SendMessageOpts{ParseMode: "html"})
				return err
			}
			_, err = b.SendAnimation(ctx.EffectiveChat.Id, pic.URL, &gotgbot.SendAnimationOpts{})
			if err != nil {
				return err
			}
		}
	}

	if !sat {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("<code>%s</code> is not a valid endpoint\nSupported endpoints:\n <code>%s</code>", args[1], nekosAPI.EndPoints), &gotgbot.SendMessageOpts{ParseMode: "html"})
		return err
	}

	return nil
}
