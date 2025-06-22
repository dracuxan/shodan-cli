package shodan

import (
	"encoding/json"
	"net/http"
)

type APIInfo struct {
	ScanCredits  int    `json:"scan_credits"`
	Plan         string `json:"plan"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
	QueryCredits int    `json:"query_credits"`
	MonitoredIps int    `json:"monitored_ips"`
	UnlockedLeft int    `json:"unlocked_left"`
	Telnet       bool   `json:"telnet"`
}

func (c *Client) APIInfo() (*APIInfo, error) {
	infoUrl := BaseUrl + "/api-info?key=" + c.apiKey
	res, err := http.Get(infoUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var info APIInfo
	if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
		return nil, err
	}

	return &info, nil
}
