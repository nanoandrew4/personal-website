package nsteg

import (
	"context"
	"personal-website/internal/logging"
)

const (
	decodeOperationPath = "/api/v1/image/decode"
)

func DecodeFilesJSON(ctx context.Context, decodeImageRequest *DecodeImageRequest, logger *logging.Logger) (*DecodeImageResponse, error) {
	var result DecodeImageResponse

	_, err := nstegClient.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(decodeImageRequest).
		SetResult(&result).
		Post(decodeOperationPath)

	if err != nil {
		logger.WithError(err).Error("Error during request to decode image with nsteg")
		return nil, err
	}

	logger.Info("Successful image decode")

	return &result, nil
}

type DecodeImageRequest struct {
	ImageToDecode []byte `json:"image_to_decode"`
}

type File struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}

type DecodeImageResponse struct {
	DecodedFiles []*File `json:"decoded_files"`
}
