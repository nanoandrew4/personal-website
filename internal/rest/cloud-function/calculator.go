package cloudfunction

import (
	"context"
	"github.com/go-resty/resty/v2"
	"os"
	"personal-website/internal/logging"
	"personal-website/internal/rest"
)

const (
	operationPath = "calculate"
)

var calculatorClient = rest.NewRestClient(resty.New().SetBaseURL(os.Getenv("CALCULATOR_BASE_URL")))

type operationRequest struct {
	Operation string `json:"operation"`
}

type OperationResponse struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

func CalculateOperation(ctx context.Context, operation string, logger *logging.Logger) (*OperationResponse, error) {
	var result OperationResponse

	response, err := calculatorClient.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(operationRequest{Operation: operation}).
		SetResult(&result).
		Post(operationPath)

	if err != nil {
		logger.WithError(err).With("operation", operation).Error("Error during request to calculate operation")
		return &result, err
	}

	logger.With("response", string(response.Body())).Info("Successful calculation")

	return &result, nil
}
