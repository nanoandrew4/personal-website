package internal

import (
	"context"
	"personal-website/constants"
	"personal-website/locales_generated"
)

var i18n *locales_generated.Localizer

func init() {
	i18n = locales_generated.New(constants.DefaultLocale, constants.DefaultLocale)
}

func Localize(ctx context.Context, key string) string {
	return i18n.GetWithLocale(ctx.Value("locale").(string), key)
}
