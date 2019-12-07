package microscope

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type client struct {
	endpoint string
}

//NewClient function creating API client
func NewClient(endpoint string) *client {
	return &client{endpoint: endpoint}
}

func (c *client) getApi(apiPath string) (*string, error) {
	resp, err := http.Get(strings.Join([]string{c.endpoint, apiPath}, "/"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	body := string(byteBody)

	return &body, nil
}

func (c *client) postApi(apiPath string, data string) (*string, error) {

	resp, err := http.Post(strings.Join([]string{c.endpoint, apiPath}, "/"), "application/json", strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	body := string(byteBody)

	return &body, nil
}
