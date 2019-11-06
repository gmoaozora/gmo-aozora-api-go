package libs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Request an URL
func Request(method string, url string, values url.Values, authHeader string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create http Request.")
	}

	if method == "post" {
		req.Method = http.MethodPost
		req.Body = ioutil.NopCloser(strings.NewReader(values.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	rClient := http.Client{
		Timeout: time.Second * 2, // timout for the request
	}

	res, err := rClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Http Client Error")
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body.")
	}

	return resBody, err
}

// FormatRequest generates ascii representation of a request
func FormatRequest(r *http.Request) string {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n--------------------\n")
	fmt.Printf(" POST REQUEST \n")
	fmt.Printf("--------------------\n")
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s = %v\n", key, value)

	}
	fmt.Printf("--------------------\n\n")
	return string(requestDump)
}
