package nsteg

import (
	"github.com/go-resty/resty/v2"
	"os"
	"personal-website/internal/rest"
)

var nstegClient = rest.NewRestClient(resty.New().SetBaseURL(os.Getenv("NSTEG_BASE_URL")))
