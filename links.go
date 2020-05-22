package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	unlock = "%s/unlock?agent=%s&apikey=%s&link=%s"
)

//UnlockLinkResponse is the response of the unlock link call
type UnlockLinkResponse struct {
	Status string `json:"status"`
	Data   struct {
		Link      string        `json:"link"`
		Host      string        `json:"host"`
		Filename  string        `json:"filename"`
		Streaming []interface{} `json:"streaming"`
		Paws      bool          `json:"paws"`
		Filesize  int           `json:"filesize"`
		Streams   []struct {
			ID       string  `json:"id"`
			Ext      string  `json:"ext"`
			Quality  string  `json:"quality"`
			Filesize int     `json:"filesize"`
			Proto    string  `json:"proto"`
			Name     string  `json:"name"`
			Tbr      float64 `json:"tbr,omitempty"`
			Abr      int     `json:"abr,omitempty"`
		} `json:"streams"`
		ID string `json:"id"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

// UnlockLink returns a downloadable link
func (c *Client) UnlockLink(link string) (UnlockLinkResponse, error) {
	resp, err := http.Get(fmt.Sprintf(unlock, magnetURL, c.AppName, c.APIKEY, link))

	if err != nil {
		return UnlockLinkResponse{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var unlockResponse UnlockLinkResponse

	err = decoder.Decode(&unlockResponse)
	if err != nil {
		return UnlockLinkResponse{}, err
	}

	if unlockResponse.Status != "success" {
		return UnlockLinkResponse{}, errors.New(unlockResponse.Error.Message)
	}

	return unlockResponse, nil
}
