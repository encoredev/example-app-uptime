package monitor

import (
	"context"
	"net/http"
	"strings"
	"time"
)

// PingResponse is the response from the Ping endpoint.
type PingResponse struct {
	Up bool `json:"up"`

	// TLSExpiry specifies when the TLS certificate expires,
	// or nil if there is no TLS certificate or if the site
	// couldn't be reached (Up == false).
	TLSExpiry *time.Time `json:"tls_expiry"`
}

// Ping pings a specific site and determines whether it's up or down right now.
//
//encore:api public path=/ping/*url
func Ping(ctx context.Context, url string) (*PingResponse, error) {
	if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
		url = "https://" + url
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &PingResponse{Up: false}, nil
	}
	defer resp.Body.Close()

	exp := getTLSExpiry(resp)
	if resp.StatusCode >= 400 {
		return &PingResponse{Up: false, TLSExpiry: exp}, nil
	}
	return &PingResponse{Up: true, TLSExpiry: exp}, nil
}

func getTLSExpiry(resp *http.Response) *time.Time {
	if resp.TLS != nil {
		return &resp.TLS.PeerCertificates[0].NotAfter
	}
	return nil
}
