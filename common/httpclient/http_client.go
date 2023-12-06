// http_client.go in github.com/unigrid-project/cosmos-sdk-common

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
	Client = cleanhttp.DefaultPooledClient()
	// Customize the Transport to disable TLS verification
	transport := Client.Transport.(*http.Transport)
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	transport.MaxIdleConnsPerHost = 50 // Set a limit for idle connections per host

	// Setting a timeout for the client
	Client.Timeout = 15 * time.Second
}
