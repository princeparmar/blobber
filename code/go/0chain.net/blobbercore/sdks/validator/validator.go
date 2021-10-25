package sdks

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/0chain/blobber/code/go/0chain.net/core/pb"
	"github.com/0chain/gosdk/core/resty"
)

type Validator struct {
}

// CreateTransport create http.Transport with default dial timeout
func (v *Validator) CreateTransport() *http.Transport {
	return &http.Transport{
		Dial: (&net.Dialer{
			Timeout: resty.DefaultDialTimeout,
		}).Dial,
		TLSHandshakeTimeout: resty.DefaultDialTimeout,
	}
}

// BuildUrls build full request url
func (v *Validator) BuildUrls(baseURLs []string, queryString map[string]string, pathFormat string, pathArgs ...interface{}) []string {

	requestURL := pathFormat
	if len(pathArgs) > 0 {
		requestURL = fmt.Sprintf(pathFormat, pathArgs...)
	}

	if len(queryString) > 0 {
		requestQuery := make(url.Values)
		for k, v := range queryString {
			requestQuery.Add(k, v)
		}

		requestURL += "?" + requestQuery.Encode()
	}

	list := make([]string, len(baseURLs))
	for k, v := range baseURLs {
		list[k] = v + requestURL
	}

	return list
}

func (v *Validator) DoPost(req *Request, handle resty.Handle) *resty.Resty {

	opts := make([]resty.Option, 0)

	opts = append(opts, resty.WithRetry(resty.DefaultRetry))
	opts = append(opts, resty.WithTimeout(resty.DefaultRequestTimeout))
	// opts = append(opts, resty.WithBefore(func(r *http.Request) {
	// 	z.SignRequest(r, req.AllocationID) //nolint
	// }))

	if len(req.ContentType) > 0 {
		opts = append(opts, resty.WithHeader(map[string]string{
			"Content-Type": req.ContentType,
		}))
	}

	r := resty.New(v.CreateTransport(), handle, opts...)

	return r
}

func (v *Validator) SendChallenge(ctx context.Context, task *pb.ChallengeTask) error {
	return nil
}
