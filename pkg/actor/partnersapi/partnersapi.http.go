package partnersapi

import (
	"bytes"
	"net/http"
)

type partnersapi struct {
	uri string
}

func NewPartnersapi(uri string) *partnersapi {
	return &partnersapi{
		uri: uri,
	}
}

func (r *partnersapi) PostApplication(b []byte) (*http.Response, error) {
	return http.Post(r.uri+"/api/applications", "application/json", bytes.NewReader(b))
}
