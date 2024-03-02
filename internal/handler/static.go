package handler

import (
	"bytes"
	"context"
	"github.com/a-h/templ"
	"io"
	"net/http"
	"personal-website/constants"
	"personal-website/internal/logging"
	"sync"
)

type StaticHandler struct {
	logger    *logging.Logger
	component templ.Component

	cachedPages sync.Map
}

type cacheKey struct {
	locale, theme string
}

func NewStaticHandler(component templ.Component) *StaticHandler {
	return &StaticHandler{
		component:   component,
		cachedPages: sync.Map{},
		logger:      logging.BuildLogger(),
	}
}

func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("serving request for URL", "url", r.URL.String())

	locale := getAndAppendLocaleToRequestContext(r)

	if page, exists := h.cachedPages.Load(cacheKey{locale: locale, theme: "default"}); exists {
		h.logger.Debug("serving page from cache")
		h.servePage(w, page.([]byte))
	} else {
		r = r.WithContext(context.WithValue(r.Context(), "locale", locale))
		http.SetCookie(w, &http.Cookie{Name: "locale", Value: locale, Path: "/"})

		byteWriter := bytes.Buffer{}
		renderErr := h.component.Render(r.Context(), io.Writer(&byteWriter))
		if renderErr != nil {
			h.logger.Error("error rendering page")
		}
		renderedBytes := byteWriter.Bytes()
		h.cachedPages.Store(cacheKey{locale: locale, theme: "default"}, renderedBytes)
		h.servePage(w, renderedBytes)
	}

	h.logger.Debug("serving response")
}

func (h *StaticHandler) servePage(w http.ResponseWriter, page []byte) {
	_, err := w.Write(page)
	if err != nil {
		h.logger.With("error", err.Error()).Error("Error serving page")
		http.Error(w, "Error serving page", http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "text/html")
	return
}

func isValidLocale(locale string) bool {
	for _, validLocale := range constants.SupportedLocales {
		if locale == validLocale {
			return true
		}
	}
	return false
}
