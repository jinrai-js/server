package jinrai

import "log"

func (c *Jinrai) Log(text ...any) {
	if !c.Verbose {
		return
	}

	log.Println(text...)
}
