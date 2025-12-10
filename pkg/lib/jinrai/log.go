package jinrai

import "log"

func (c *Jinrai) Log(text ...any) {
	if !c.Server.Verbose {
		return
	}

	log.Println(text...)
}
