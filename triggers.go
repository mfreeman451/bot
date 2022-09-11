package main

import (
	"fmt"
	hbot "github.com/whyrusleeping/hellabot"
)

var SigPic = hbot.Trigger{
	// Condition
	func(b *hbot.Bot, m *hbot.Message) bool {
		return m.Content == "penis"
		// return m.From == "whyrusleeping"
	},
	// Action
	func(b *hbot.Bot, m *hbot.Message) bool {
		msg := fmt.Sprintf("%s said %s", m.From, m.Content)
		b.Msg("#Chases", msg)
		// b.Reply(m, "whyrusleeping said something")
		return false
	},
}
