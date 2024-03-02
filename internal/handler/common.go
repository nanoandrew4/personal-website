package handler

import (
	"context"
	"net/http"
	"personal-website/constants"
)

func getAndAppendLocaleToRequestContext(r *http.Request) string {
	var locale string
	localeCookie, err := r.Cookie("locale")
	if err != nil || !isValidLocale(localeCookie.Value) {
		locale = constants.DefaultLocale
	} else {
		locale = localeCookie.Value
	}

	*r = *r.WithContext(context.WithValue(r.Context(), "locale", locale))
	//http.SetCookie(w, &http.Cookie{Name: "locale", Value: locale, Path: "/"})

	return locale
}
