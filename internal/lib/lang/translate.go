package lang

import (
	"context"
	"net/http"

	"github.com/jinrai-js/server/internal/lib/config"
	"github.com/jinrai-js/server/internal/lib/lang/lang_base"
	"github.com/jinrai-js/server/internal/lib/lang/lang_context"
)

func Translate(ctx context.Context, key string) string {
	lang := lang_context.Get(ctx)

	if lang.Active == lang.Default {
		return key
	}

	return lang_base.GetValue(ctx, lang.SourceUrl, lang.Active, key)
}

func CreateLang(c *config.JsonConfig, r *http.Request) lang_context.Lang {
	result := lang_context.Lang{
		SourceUrl: c.Lang.LangBaseURL,
		Default:   c.Lang.DefaultLang,
		Active:    c.Lang.DefaultLang,
	}

	if c.Lang.Source.From == "cookie" {
		if cookie, err := r.Cookie(c.Lang.Source.Key); err == nil {
			result.Active = cookie.Value
		}
	}

	return result
}
