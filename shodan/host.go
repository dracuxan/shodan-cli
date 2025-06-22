package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShodanData struct {
	RegionCode   *string    `json:"region_code"`
	IP           int        `json:"ip"`
	PostalCode   *string    `json:"postal_code"`
	CountryCode  string     `json:"country_code"`
	City         *string    `json:"city"`
	DMACode      *string    `json:"dma_code"`
	LastUpdate   string     `json:"last_update"`
	Latitude     float64    `json:"latitude"`
	Tags         []string   `json:"tags"`
	AreaCode     *string    `json:"area_code"`
	CountryName  string     `json:"country_name"`
	Hostnames    []string   `json:"hostnames"`
	Org          string     `json:"org"`
	Data         []DataItem `json:"data"`
	ASN          string     `json:"asn"`
	ISP          string     `json:"isp"`
	Longitude    float64    `json:"longitude"`
	CountryCode3 *string    `json:"country_code3"`
	Domains      []string   `json:"domains"`
	IPStr        string     `json:"ip_str"`
	OS           *string    `json:"os"`
	Ports        []int      `json:"ports"`
}

type DataItem struct {
	Shodan    ShodanInfo `json:"_shodan"`
	Hash      int        `json:"hash"`
	OS        *string    `json:"os"`
	Opts      Opts       `json:"opts"`
	IP        int        `json:"ip"`
	ISP       string     `json:"isp"`
	Port      int        `json:"port"`
	Hostnames []string   `json:"hostnames"`
	Location  Location   `json:"location"`
	DNS       DNSInfo    `json:"dns"`
	Timestamp string     `json:"timestamp"`
	Domains   []string   `json:"domains"`
	Org       string     `json:"org"`
	Data      string     `json:"data"`
	ASN       string     `json:"asn"`
	Transport string     `json:"transport"`
	IPStr     string     `json:"ip_str"`
}

type ShodanInfo struct {
	ID      string         `json:"id"`
	Options map[string]any `json:"options"`
	PTR     bool           `json:"ptr"`
	Module  string         `json:"module"`
	Crawler string         `json:"crawler"`
}

type Opts struct {
	Raw string `json:"raw"`
}

type Location struct {
	City         *string `json:"city"`
	RegionCode   *string `json:"region_code"`
	AreaCode     *string `json:"area_code"`
	Longitude    float64 `json:"longitude"`
	CountryCode3 *string `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	PostalCode   *string `json:"postal_code"`
	DMACode      *string `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	Latitude     float64 `json:"latitude"`
}

type DNSInfo struct {
	ResolverHostname *string `json:"resolver_hostname"`
	Recursive        bool    `json:"recursive"`
	ResolverID       *string `json:"resolver_id"`
	Software         *string `json:"software"`
}

func (c *Client) HostInfo(q string) (*ShodanData, error) {
	searchUrl := fmt.Sprintf("%s/shodan/host/%s?key=%s", BaseUrl, q, c.apiKey)
	res, err := http.Get(searchUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var searchResults ShodanData
	if err := json.NewDecoder(res.Body).Decode(&searchResults); err != nil {
		return nil, err
	}
	return &searchResults, nil
}
