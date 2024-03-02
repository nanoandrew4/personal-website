package nsteg

import (
	"encoding/base64"
	"fmt"
	"github.com/a-h/templ"
	"io"
	"net/http"
	"personal-website/internal/logging"
	"personal-website/internal/rest/api/nsteg"
	"strconv"
)

func EncodeImage(r *http.Request, logger *logging.Logger) templ.Component {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return encodingError("projects.nsteg.generic-error")
	}

	imageEncodeRequestBody, imageName, errFrag := getImageToEncodeBytes(r)
	if errFrag != nil {
		return errFrag
	}

	imageEncodeRequestBody.FilesToHide, errFrag = getFilesToHide(r)
	if errFrag != nil {
		return errFrag
	}

	// TODO: validate sizes before forwarding requests to try to catch some errors early
	// TODO: release memory ASAP
	response, err := nsteg.EncodeImageAsJSON(r.Context(), imageEncodeRequestBody, logger)
	if err != nil {
		return encodingError("projects.nsteg.encode-error")
	}

	base64Img := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(response.EncodedImage))

	return encodedImageFragment(fmt.Sprintf("encoded-%s", imageName), base64Img)
}

func getImageToEncodeBytes(r *http.Request) (*nsteg.EncodeImageRequest, string, templ.Component) {
	lsbsToUseStr := r.FormValue("lsbsToUse")
	lsbsToUse, err := strconv.ParseInt(lsbsToUseStr, 10, 64)
	if err != nil {
		return nil, "", encodingError("projects.nsteg.invalid-lsbs")
	}

	imageToEncode, imageHeader, err := r.FormFile("imageToEncode")
	if err != nil {
		return nil, "", encodingError("projects.nsteg.image-invalid")
	}
	defer imageToEncode.Close()

	imageBytes, err := io.ReadAll(imageToEncode)
	if err != nil {
		return nil, "", encodingError("projects.nsteg.image-invalid")
	}

	return &nsteg.EncodeImageRequest{
		ImageToEncode: imageBytes,
		LsbsToUse:     byte(lsbsToUse),
	}, imageHeader.Filename, nil
}

func getFilesToHide(r *http.Request) ([]nsteg.FileToHide, templ.Component) {
	var filesToHideBytes []nsteg.FileToHide

	for _, fileHeader := range r.MultipartForm.File["filesToHide"] {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, encodingError("projects.nsteg.files-to-hide-invalid")
		}

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return nil, encodingError("projects.nsteg.files-to-hide-invalid")
		}
		filesToHideBytes = append(filesToHideBytes, nsteg.FileToHide{
			Name:    fileHeader.Filename,
			Content: fileBytes,
		})
	}

	return filesToHideBytes, nil
}
