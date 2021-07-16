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
			mimeType := getType(pic.URL)
			if mimeType == "image/jpeg" {
				_, err = b.SendPhoto(ctx.EffectiveChat.Id, pic.URL, &gotgbot.SendPhotoOpts{})
			} else if mimeType == "image/gif" {
				_, err = b.SendAnimation(ctx.EffectiveChat.Id, pic.URL, &gotgbot.SendAnimationOpts{})
			} else {
				_, err = b.SendMessage(ctx.EffectiveChat.Id, fmt.Sprintf("Type <code>%s</code> not yet handled", mimeType), &gotgbot.SendMessageOpts{ParseMode: "html"})
			}
			if err != nil {
				return err
			}
		}
	}

	if !sat {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("<code>%s</code> is not a valid endpoint\n\n<b>Supported endpoints</b>:\n%s", args[1], nekosAPI.GetEndpointsHTML()), &gotgbot.SendMessageOpts{ParseMode: "html"})
		return err
	}

	return nil
}
