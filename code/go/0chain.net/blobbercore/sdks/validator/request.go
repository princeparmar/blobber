package sdks

import "bytes"

// Request request payload
type Request struct {
	// ContentType content-type in header
	ContentType string
	// Body form data
	Body *bytes.Buffer
	// QueryString query string
	QueryString map[string]string
}
