package partnersapi

import "net/http"

type Partnersapi interface {
	PostApplication(b []byte) (*http.Response, error)
}
