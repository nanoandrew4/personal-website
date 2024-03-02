package handler

import (
	"github.com/a-h/templ"
	"net/http"
	"personal-website/internal/logging"
)

type PostRequestHandler struct {
	logger *logging.Logger

	componentBuilderFunc func(r *http.Request, logger *logging.Logger) templ.Component
}

func NewPostRequestHandler(builderFunc func(*http.Request, *logging.Logger) templ.Component) *PostRequestHandler {
	return &PostRequestHandler{
		logger:               logging.BuildLogger(),
		componentBuilderFunc: builderFunc,
	}
}

func (h *PostRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = getAndAppendLocaleToRequestContext(r)

	templ.Handler(h.componentBuilderFunc(r, h.logger)).ServeHTTP(w, r)
}
