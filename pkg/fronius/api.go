package fronius

import (
	"encoding/json"
	"io/ioutil"

	api "github.com/siredmar/solis/pkg/api/froniusv1"
	"github.com/siredmar/solis/pkg/client"
)

const (
	// GetPowerFlowRealtimeDataURI is the URI for the power flow realtime data endpoint
	GetPowerFlowRealtimeDataURI = "/solar_api/v1/GetPowerFlowRealtimeData.fcgi"
)

// Status returns the status of the sonnen reader
func GetPowerFlowRealtimeData(c *client.Client) (*api.GetPowerFlowRealtimeData, error) {
	resp, err := c.Get(c.Host + GetPowerFlowRealtimeDataURI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var status *api.GetPowerFlowRealtimeData
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &status); err != nil {
		return nil, err
	}
	return status, nil
}
