package dothill

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request : Used internally, and can be used to send custom requests (see Client.Request())
type Request struct {
	Endpoint string
	Data     interface{}
}

func (req *Request) execute(client *Client) ([]byte, error) {
	url := fmt.Sprintf("%s/api%s", client.Options.Addr, req.Endpoint)
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("API returned unexpected HTTP status %d", res.StatusCode)
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
