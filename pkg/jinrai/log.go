package jinrai

import "log"

func (c *Static) Log(text ...any) {
	if !c.Verbose {
		return
	}

	log.Println(text...)
}
