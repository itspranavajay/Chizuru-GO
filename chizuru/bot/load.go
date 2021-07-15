package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

func LoadHandlers(d *ext.Dispatcher) {
	startCmd := handlers.NewCommand("start", startCMD)
	d.AddHandler(startCmd)
	d.AddHandler(handlers.NewCommand("nekos", nekosLIFE))
	d.AddHandler(handlers.NewCommand("help", getHelp))
	d.AddHandler(handlers.NewCommand("stats", getStats))
	d.AddHandler(handlers.NewMessage(filters.All, logStuff))
}
