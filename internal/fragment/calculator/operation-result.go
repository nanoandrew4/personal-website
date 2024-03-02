package calculator

import (
	"github.com/a-h/templ"
	"io"
	"net/http"
	"net/url"
	"personal-website/internal/logging"
	cloudfunction "personal-website/internal/rest/cloud-function"
	"strings"
)

func CalculateOperationResult(r *http.Request, logger *logging.Logger) templ.Component {
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		logger.WithError(readErr).Error("Error reading operation from request body")
	}

	operation, err := url.QueryUnescape(strings.Split(string(body), "=")[1])
	if err != nil {
		logger.WithError(err).Error("Error unescaping operation")
	}

	response, err := cloudfunction.CalculateOperation(r.Context(), operation, logger)
	if response != nil && response.Error != "" {
		return operationError("projects.calculator.operation-error")
	} else if err != nil {
		return operationError("projects.calculator.server-error")
	}
	return operationResult(response.Result)
}
