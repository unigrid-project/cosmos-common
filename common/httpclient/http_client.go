// http_client.go in github.com/unigrid-project/cosmos-sdk-common

package httpclient

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-cleanhttp"
)

var HttpClient *http.Client

func init() {
	// Initialize HttpClient with cleanhttp for default settings
	HttpClient = cleanhttp.DefaultPooledClient()
	// Customizing the Transport, if necessary
	transport := HttpClient.Transport.(*http.Transport)
	transport.MaxIdleConnsPerHost = 50 // Set a limit for idle connections per host

	// Setting a timeout for the client
	HttpClient.Timeout = 15 * time.Second
}
