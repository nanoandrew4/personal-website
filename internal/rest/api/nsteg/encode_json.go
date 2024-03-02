package nsteg

import (
	"context"
	"personal-website/internal/logging"
)

const (
	encodeOperationPath = "/api/v1/image/encode"
)

func EncodeImageAsJSON(ctx context.Context, encodeImageRequest *EncodeImageRequest, logger *logging.Logger) (*EncodeImageResponse, error) {
	var result EncodeImageResponse

	_, err := nstegClient.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(encodeImageRequest).
		SetResult(&result).
		Post(encodeOperationPath)

	if err != nil {
		logger.WithError(err).Error("Error during request to encode image with nsteg")
		return nil, err
	}

	logger.Info("Successful image encode")

	return &result, nil
}

type EncodeImageRequest struct {
	LsbsToUse     byte         `json:"lsbs_to_use"`
	ImageToEncode []byte       `json:"image_to_encode"`
	FilesToHide   []FileToHide `json:"files_to_hide"`
}

type FileToHide struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}

type EncodeImageResponse struct {
	EncodedImage []byte `json:"encoded_image"`
}
