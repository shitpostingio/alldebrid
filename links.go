package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	unlock    = "%s/unlock?agent=%s&apikey=%s&link=%s"
	streaming = "%s/streaming?agent=%s&apikey=%s&stream=%s&id=%s"
	delayed   = "%s/delayed?agent=%s&apikey=%s&id=%s"
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

//StreamingResponse is the response of the streaming call
type StreamingResponse struct {
	Status string `json:"status"`
	Data   struct {
		Link     string  `json:"link,omitempty"`
		Filename string  `json:"filename"`
		Filesize float64 `json:"filesize"`
		Delayed  int     `json:"delayed,omitempty"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

//DelayedResponse is the response of the delayed call
type DelayedResponse struct {
	Status string `json:"status"`
	Data   struct {
		Status   int    `json:"status"`
		Speed    int    `json:"speed"`
		TimeLeft int    `json:"time_left"`
		Progress int    `json:"progress"`
		Link     string `json:"link"`
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

//StreamingLink returns a delayed id or the direct link
func (c *Client) StreamingLink(stream, id string) (StreamingResponse, error) {
	resp, err := http.Get(fmt.Sprintf(streaming, magnetURL, c.AppName, c.APIKEY, stream, id))

	if err != nil {
		return StreamingResponse{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var streamResponse StreamingResponse

	err = decoder.Decode(&streamResponse)
	if err != nil {
		return StreamingResponse{}, err
	}

	if streamResponse.Status != "success" {
		return StreamingResponse{}, errors.New(streamResponse.Error.Message)
	}

	return streamResponse, nil
}

// DelayedLink returns a downloadable link for the given delayed id
func (c *Client) DelayedLink(delayedID string) (DelayedResponse, error) {
	resp, err := http.Get(fmt.Sprintf(delayed, magnetURL, c.AppName, c.APIKEY, delayedID))

	if err != nil {
		return DelayedResponse{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var delayedResponse DelayedResponse

	err = decoder.Decode(&delayedResponse)
	if err != nil {
		return DelayedResponse{}, err
	}

	if delayedResponse.Status != "success" {
		return DelayedResponse{}, errors.New(delayedResponse.Error.Message)
	}

	return delayedResponse, nil
}
