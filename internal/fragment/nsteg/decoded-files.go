package nsteg

import (
	"encoding/base64"
	"github.com/a-h/templ"
	"io"
	"net/http"
	"personal-website/internal/logging"
	"personal-website/internal/rest/api/nsteg"
	"strings"
)

const (
	dataURLPrefix = "data:application/octet-stream;base64,"
)

func DecodeFiles(r *http.Request, logger *logging.Logger) templ.Component {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return decodingError("projects.nsteg.generic-error")
	}

	imageDecodeRequestBody, _, errFrag := getImageToDecode(r)
	if errFrag != nil {
		return errFrag
	}

	// TODO: validate sizes before forwarding requests to try to catch some errors early
	// TODO: release memory ASAP
	response, err := nsteg.DecodeFilesJSON(r.Context(), imageDecodeRequestBody, logger)
	if err != nil {
		return decodingError("projects.nsteg.encode-error")
	}

	for _, decodedFile := range response.DecodedFiles {
		var hrefBuilder strings.Builder
		hrefBuilder.Grow(len(decodedFile.Content) + len(dataURLPrefix))
		hrefBuilder.WriteString(dataURLPrefix)
		hrefBuilder.WriteString(base64.StdEncoding.EncodeToString(decodedFile.Content))
		decodedFile.Content = []byte(hrefBuilder.String())
	}

	return decodedFiles(response.DecodedFiles)
}

func getImageToDecode(r *http.Request) (*nsteg.DecodeImageRequest, string, templ.Component) {
	imageToDecode, imageHeader, err := r.FormFile("imageToDecode")
	if err != nil {
		return nil, "", decodingError("projects.nsteg.image-invalid")
	}
	defer imageToDecode.Close()

	imageBytes, err := io.ReadAll(imageToDecode)
	if err != nil {
		return nil, "", decodingError("projects.nsteg.image-invalid")
	}

	return &nsteg.DecodeImageRequest{
		ImageToDecode: imageBytes,
	}, imageHeader.Filename, nil
}
