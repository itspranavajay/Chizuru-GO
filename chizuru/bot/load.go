package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadHandlers(d *ext.Dispatcher) {
	startCmd := handlers.NewCommand("start", startCMD)
	d.AddHandler(startCmd)
}
