// http_client.go in github.com/unigrid-project/cosmos-common

package httpclient

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/hashicorp/go-cleanhttp"
)

var Client *http.Client

func init() {
	// Initialize Client with cleanhttp for default settings
	Client = cleanhttp.DefaultClient()
	// Customize the Transport to disable TLS verification
	transport := Client.Transport.(*http.Transport)
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	transport.MaxIdleConnsPerHost = 5 // Set a limit for idle connections per host

	// Setting a timeout for the client
	Client.Timeout = 5 * time.Second
}
