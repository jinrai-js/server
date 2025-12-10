package serverState

import (
	"strings"

	"github.com/jinrai-js/go/pkg/jinrai/conventor"
)

func (i *StateInterface) getStringKey(keys []string) string {
	data := conventor.Parse(i.Key)

	switch v := data.(type) {
	case string:
		return v

	case []string:
		return strings.Join(v, "")
	}

	return ""
}
