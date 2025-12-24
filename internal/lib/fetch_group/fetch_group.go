package fetch_group

import (
	"sync"

	"github.com/jinrai-js/server/internal/lib/pass"
)

type FetchGroup struct {
	wg  sync.WaitGroup
	use bool
}

var globalFetchGroup = FetchGroup{}
var inWork = make(map[string]bool)

func Reset() {
	globalFetchGroup.use = false
}

func Run(key string) {
	if _, exists := inWork[key]; exists {
		pass.Exit()
	}

	globalFetchGroup.use = true
	globalFetchGroup.wg.Add(1)
}

func Done(key string) {
	globalFetchGroup.wg.Done()
	delete(inWork, key)
}

func Wait() {
	globalFetchGroup.wg.Wait()
}

func WasSandRequest() bool {
	return globalFetchGroup.use
}
