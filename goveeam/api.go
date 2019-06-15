package goveeam

import (
	"net/http"
	"net/url"
)

type Client struct {
	APIVersion string
	VeeamEntToken string
	VeeamEntAuthHeader string
	ENTHREF url.URL
	Http http.Client
}

