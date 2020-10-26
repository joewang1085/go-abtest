package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	defaultHTTPClient *http.Client = http.DefaultClient

	port string = "9527"

	address string = "127.0.0.1"
)

func getRemoteConfig(projectID string, timeout time.Duration) ([]byte, error) {
	url := fmt.Sprintf("http://%s:%s/DescribeConfig", address, port)
	defaultHTTPClient.Timeout = timeout
	params := make(map[string][]string)
	params["project"] = []string{projectID}
	resp, err := defaultHTTPClient.PostForm(url, params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
