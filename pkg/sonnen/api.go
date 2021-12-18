package sonnen

import (
	"encoding/json"
	"io/ioutil"

	api "github.com/siredmar/solis/pkg/api/sonnenv2"
	"github.com/siredmar/solis/pkg/client"
)

const (
	// StatusURI is the URI for the status endpoint
	StatusURI = "/api/v2/status"
	// LaststatusURI is the URI for the last status endpoint
	LastStatusURI = "/api/v2/laststatus"
)

// Status returns the status of the sonnen reader
func Status(c *client.Client) (*api.Status, error) {
	resp, err := c.Get(c.Host + StatusURI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var status *api.Status
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &status); err != nil {
		return nil, err
	}
	return status, nil
}

// LastStatus returns the last status of the sonnen reader
func LastStatus(c *client.Client) (*api.LastStatus, error) {
	resp, err := c.Get(c.Host + LastStatusURI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var status *api.LastStatus
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &status); err != nil {
		return nil, err
	}
	return status, nil
}
