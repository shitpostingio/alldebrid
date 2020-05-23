package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//Domains is the domains response struct
type Domains struct {
	Status string `json:"status"`
	Data   struct {
		Hosts       []string `json:"hosts"`
		Streams     []string `json:"streams"`
		Redirectors []string `json:"redirectors"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

//GetDomainsOnly returns list of supported hosts domains and redirectors
func (c *Client) GetDomainsOnly() (Domains, error) {
	resp, err := http.Get(fmt.Sprintf(hostsdomains, hosts, c.appName))

	if err != nil {
		return Domains{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var doms Domains

	err = decoder.Decode(&doms)
	if err != nil {
		return Domains{}, err
	}

	if doms.Status != "success" {
		return Domains{}, errors.New(doms.Error.Message)
	}

	return doms, nil
}
